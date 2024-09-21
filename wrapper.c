#include "wrapper.h"
#include <stdio.h>
#include <sys/ptrace.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>
#include <sys/reg.h>
#include <sys/user.h>

char* getString(pid_t pid, long addr, size_t size);

int wrapper_connect(pid_t pid) {
    return ptrace(PTRACE_ATTACH, pid, NULL, NULL);
}

int wrapper_waitpid(pid_t pid, int* status) {
    waitpid(pid, status, 0);
    return 0;
}

int wrapper_detach(int pid) {
    //printf("detaching from pid: %d\n", pid);
    int ret = ptrace(PTRACE_DETACH, pid, NULL, NULL);
    //printf("ptrace returned: %d\n", ret);
    //fflush(stdout);

    return ret;
}

int wrapper_get_sys_call(pid_t pid) {
    //printf("peeking to attached pid: %d\n", pid);
    return ptrace(PTRACE_PEEKUSER, pid, sizeof(long) * ORIG_RAX, NULL);
}

int wrapper_has_exited(int status) {
    return WIFEXITED(status);
}

void wrapper_continue(pid_t pid) {
    ptrace(PTRACE_SYSCALL, pid, NULL, NULL);
}

char* wrapper_peek_regs(pid_t pid) {
    struct user_regs_struct regs;
    ptrace(PTRACE_GETREGS, pid, NULL, &regs);
    //printf("Write called with ""%llu, %llu, %llu\n", regs.rdi, regs.rsi, regs.rdx);
    return getString(pid, regs.rsi, regs.rdx);
}

char* getString(pid_t pid, long addr, size_t size) {
    char *buffer = (char*)malloc(size + 1);
    if (!buffer) {
        perror("malloc failed");
        return NULL;
    }

    // Read the string data from the target process
    for (size_t i = 0; i < size; i += sizeof(long)) {
        long data = ptrace(PTRACE_PEEKDATA, pid, addr + i, NULL);
        if (data == -1 && errno) {
            perror("ptrace PEEKDATA failed");
            free(buffer);
            return NULL;
        }
        memcpy(buffer + i, &data, sizeof(long));
    }

    buffer[size] = '\0'; // Null-terminate the string
    return buffer;
}
