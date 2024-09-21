#ifndef WRAPPER_H
#define WRAPPER_H

#include <sys/wait.h>

int wrapper_get_sys_call(pid_t pid);
int wrapper_detach(pid_t pid);
int wrapper_waitpid(pid_t pid, int* status);
int wrapper_connect(pid_t pid);
int wrapper_has_exited(int status);
void wrapper_continue(pid_t pid);
char* wrapper_peek_regs(pid_t pid);

#endif
