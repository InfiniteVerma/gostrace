package main

import (
    "fmt"
    "os"
    "strconv"
)

/*
#include <sys/ptrace.h>
#include <sys/wait.h>
#include <stdio.h>
int hello(int pid) {
    printf("Hello from C\n");
    fflush(stdout);
    int ret = ptrace(PTRACE_ATTACH, pid, NULL, NULL);
    printf("%d\n", ret);
    fflush(stdout);

    waitpid(pid, NULL, 0); // wait for process to stop

    while (1) {
        if (ptrace(PTRACE_SYSCALL, pid, NULL, NULL) == -1) {
            perror("ptrace(PTRACE_SYSCALL)");
            return 1;
        }
        waitpid(pid, NULL, 0); // Wait for syscall entry
        // Read system call number and arguments here
        if (ptrace(PTRACE_SYSCALL, pid, NULL, NULL) == -1) {
            perror("ptrace(PTRACE_SYSCALL)");
            return 1;
        }
        waitpid(pid, NULL, 0); // Wait for syscall exit
        // Read syscall return value here
    }
}
*/
import "C"

func main() {
    
    args_length := len(os.Args)

    if args_length != 2 {
        fmt.Println("ERROR invalid number of args passed. Usage: ./goptrace <pid>")
        os.Exit(1)
    }

    pid := os.Args[1]

    fmt.Println("Starting ptrace on pid:", pid)

    i, err := strconv.Atoi(pid)

    if err != nil {
        fmt.Println("ERROR couldn't cast string pid to int")
        os.Exit(1)
    }

    ret := C.hello(C.int(i));
    //res := int(C.ptrace());
}
