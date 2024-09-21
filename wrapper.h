#ifndef WRAPPER_H
#define WRAPPER_H

#include <sys/wait.h>

int ptrace_get_sys_call(int pid);
int ptrace_detach(int pid);
int c_waitpid(pid_t pid);
int ptrace_connect(int pid);

#endif
