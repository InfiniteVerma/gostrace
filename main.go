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

type StraceParams struct {
}

func searchStr(slice []string, str string) (bool, int) {
	for i, v := range slice {
		if v == str {
			return true, i
		}
	}
	return false, -1
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

	if found, _ := searchStr(slice, "-h"); found {
		fmt.Println(USAGE)
		os.Exit(1)
	}

	if found, index := searchStr(slice, "-p"); found {
		if index+2 >= len(os.Args) {
			fmt.Println("ERROR couldn't find pid.")
			os.Exit(1)
		}

		pid_str := os.Args[index+2]
		fmt.Println("Starting gostrace with pid: ", pid_str)
		i, err := strconv.Atoi(pid_str)

		if err != nil {
			fmt.Println("ERROR couldn't cast string pid to int")
			os.Exit(1)
		}

		return i
	}

	fmt.Println(HELP)
	os.Exit(1)
	return 1
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

	C.wrapper_hello()
}
