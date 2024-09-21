package main

import (
    "fmt"
    "os"
    "strconv"
)

/*
#cgo CFLAGS: -I.
#include "wrapper.h"
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

    C.wrapper_hello()
}
