#include <stdio.h>
#include <stdlib.h>
#include <sys/ptrace.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>
#include <errno.h>
#include <string.h>
#include <sys/reg.h>
#include <sys/syscall.h>

void print_string(pid_t target_pid, long addr, size_t size) {
    char *buffer = (char*)malloc(size + 1);
    if (!buffer) {
        perror("malloc failed");
        return;
    }

    // Read the string data from the target process
    for (size_t i = 0; i < size; i += sizeof(long)) {
        long data = ptrace(PTRACE_PEEKDATA, target_pid, addr + i, NULL);
        if (data == -1 && errno) {
            perror("ptrace PEEKDATA failed");
            free(buffer);
            return;
        }
        memcpy(buffer + i, &data, sizeof(long));
    }

    buffer[size] = '\0'; // Null-terminate the string
    printf("String being written: %s\n", buffer);
    free(buffer);
}

int main() {
    pid_t target_pid = 245312; // Replace with your target process PID

    // Attach to the process
    if (ptrace(PTRACE_ATTACH, target_pid, NULL, NULL) == -1) {
        perror("ptrace attach failed");
        return 1;
    }

    // Allow the target process to run
    ptrace(PTRACE_CONT, target_pid, NULL, NULL);

    while (1) {
        // Wait for the target process to stop
        int status;
        waitpid(target_pid, &status, 0);

        // Check if the target process has exited
        if (WIFEXITED(status)) {
            printf("Target process exited\n");
            break;
        }

        // Inspect the process to check if it made a syscall
        long orig_eax = ptrace(PTRACE_PEEKUSER, target_pid, sizeof(long) * ORIG_RAX, NULL);
        if (orig_eax == -1) {
            perror("ptrace PEEKUSER failed");
            break;
        }

        // If we hit a syscall
        if (orig_eax == SYS_write) {
            printf("Found sys_write\n");

            long rdi_val = ptrace(PTRACE_PEEKUSER, target_pid, sizeof(long) * RDI, NULL);
            long rsi_val = ptrace(PTRACE_PEEKUSER, target_pid, sizeof(long) * RSI, NULL);
            long rdx_val = ptrace(PTRACE_PEEKUSER, target_pid, sizeof(long) * RDX, NULL);

            printf("File descriptor (rdi): %ld\n", rdi_val);
            printf("Data address (rsi): %ld\n", rsi_val);
            printf("Data size (rdx): %ld\n", rdx_val);

            // Print the string being written
            print_string(target_pid, rsi_val, rdx_val);
        }

        // Continue the process to the next syscall
        ptrace(PTRACE_SYSCALL, target_pid, NULL, NULL);
    }

    // Detach from the process when done
    if (ptrace(PTRACE_DETACH, target_pid, NULL, NULL) == -1) {
        perror("ptrace detach failed");
        return 1;
    }

    return 0;
}
