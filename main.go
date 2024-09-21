package main

import (
    "fmt"
    "os"
    "strconv"
)

/*
#include <sys/ptrace.h>
#include <sys/reg.h>
#include <sys/wait.h>
#include <stdio.h>

int ptrace_connect(int pid) {
    printf("hello world, connecting to pid: %d\n", pid);
    int ret = ptrace(PTRACE_ATTACH, pid, NULL, NULL);
    printf("ptrace returned: %d\n", ret);
    fflush(stdout);

    return ret;
}

int c_waitpid(pid_t pid) {
    waitpid(pid, NULL, 0); // wait for process to stop
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

    ret := C.ptrace_connect(C.int(i))

    if ret < 0 {
        fmt.Println("ERROR could not connect to pid: ", pid)
        os.Exit(1)
    }

    fmt.Println("ptrace_connect returned ret: ", ret)
    fmt.Println("now waiting")

    C.c_waitpid(C.int(i))

    fmt.Println("\n\nLooking at orig_eax: ")

    sys_int := int(C.ptrace_get_sys_call(C.int(i)))

    fmt.Println("sys call called: ", get_sys_call(sys_int))

    ret = C.ptrace_detach(C.int(i));

    if ret < 0 {
        fmt.Println("ERROR could not connect to pid: ", pid)
        os.Exit(1)
    }
}
