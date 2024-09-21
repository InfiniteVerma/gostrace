#include <sys/ptrace.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/reg.h>
#include <sys/wait.h>

int main() {
    pid_t target_pid; // Replace with your target process PID
    printf("Enter PID to attach: ");
    scanf("%d", &target_pid);

    // Attach to the process
    if (ptrace(PTRACE_ATTACH, target_pid, NULL, NULL) == -1) {
        perror("ptrace attach failed");
        return 1;
    }

    // Wait for the target process to stop
    waitpid(target_pid, NULL, 0);

    // You can now inspect or manipulate the process here
    printf("peeking to attached pid: %d\n", target_pid);
    int orig_eax = ptrace(PTRACE_PEEKUSER, target_pid, sizeof(long) * ORIG_RAX, NULL);
    printf("orig_eax val: %d\n", orig_eax);
    fflush(stdout);

    sleep(20);

    // Detach from the process when done
    //if (ptrace(PTRACE_DETACH, target_pid, NULL, NULL) == -1) {
    //    perror("ptrace detach failed");
    //    return 1;
    //}

    return 0;
}

