#include "wrapper.h"
#include <stdio.h>
#include <sys/ptrace.h>
#include <sys/reg.h>

int wrapper_hello() {
    printf("Hello from wrapper!\n");
    fflush(stdout);
    return 0;
}

int ptrace_connect(int pid) {
    printf("hello world, connecting to pid: %d\n", pid);
    int ret = ptrace(PTRACE_ATTACH, pid, NULL, NULL);
    printf("ptrace returned: %d\n", ret);
    fflush(stdout);

    return ret;
}

int c_waitpid(pid_t pid) {
    waitpid(pid, NULL, 0); // wait for process to stop
    return 0;
}

int ptrace_detach(int pid) {
    printf("detaching from pid: %d\n", pid);
    int ret = ptrace(PTRACE_DETACH, pid, NULL, NULL);
    printf("ptrace returned: %d\n", ret);
    fflush(stdout);

    return ret;
}

int ptrace_get_sys_call(int pid) {
    printf("peeking to attached pid: %d\n", pid);
    ptrace(PTRACE_PEEKUSER, pid, sizeof(long) * ORIG_RAX, NULL);
}

