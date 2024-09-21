package main

/*
#cgo CFLAGS: -I.
#include "wrapper.h"
*/
import "C"

import (
	"fmt"
	"os"
	"strconv"
)

type GoStraceParams struct {
    helpSet bool
    pidSet bool
    pid int
    err bool
}

const HELP = "gostrace: must have PROG [ARGS] or -p PID. Try 'gostrace -h' for more information."
const USAGE = "Usage: gostrace [-p PID]"

func parseArgs() int {

	args_length := len(os.Args)

	fmt.Println("args passed: ", args_length)

	if args_length == 1 {
		fmt.Println(HELP)
		os.Exit(1)
	}

	slice := os.Args[1:]

    param := GoStraceParams{}
    //for i, v := range slice {
    for i:=0; i < len(slice); i++ {
        v := slice[i]
        fmt.Println("i: ", i, "v: ", v)

        if v == "-h" {
            param.helpSet = true
        } else if v == "-p" {
            param.pidSet = true

            if i + 1 >= args_length-1 {
                fmt.Println("ERROR couldn't find pid.")
                os.Exit(1)
            }

            pid_str := slice[i + 1]
	        pid_int, err := strconv.Atoi(pid_str)

            if err != nil {
                fmt.Println("ERROR could not parse pid to int ", err)
                os.Exit(1)
            }

            param.pid = pid_int
            i++
        } else {
            param.err = true
        }
    }

    fmt.Println(param)

    if param.err {
        fmt.Println("ERROR args parsing failed.")
        os.Exit(1)
    }

    if param.helpSet {
		fmt.Println(USAGE)
		os.Exit(1)
    }

    if param.pidSet != true {
        fmt.Println("ERROR pid was not passed. Use -h to get usage")
        os.Exit(1)
    }

	fmt.Println("Starting gostrace with pid: ", param.pid)
	return param.pid
}

func main() {

	pid := parseArgs()

	fmt.Println("Starting ptrace on pid:", pid)

	ret := C.ptrace_connect(C.int(pid))

	if ret < 0 {
		fmt.Println("ERROR could not connect to pid: ", pid)
		os.Exit(1)
	}

	fmt.Println("ptrace_connect returned ret: ", ret)
	fmt.Println("now waiting")

	C.c_waitpid(C.int(pid))

	fmt.Println("\n\nLooking at orig_eax: ")

	sys_int := int(C.ptrace_get_sys_call(C.int(pid)))

	fmt.Println("sys call called: ", get_sys_call(sys_int))

	ret = C.ptrace_detach(C.int(pid))

	if ret < 0 {
		fmt.Println("ERROR could not connect to pid: ", pid)
		os.Exit(1)
	}
}
