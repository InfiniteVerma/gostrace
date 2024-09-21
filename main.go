package main

/*
#cgo CFLAGS: -I.
#include <stdlib.h>
#include "wrapper.h"
*/
import "C"

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

type GoStraceParams struct {
	helpSet bool
	pidSet  bool
	pid     int
	err     bool
}

type UserRegsStruct struct {
    r15 uint64
    r14 uint64
    r13 uint64
    r12 uint64
    rbp uint64
    rbx uint64
    r11 uint64
    r10 uint64
    r9 uint64
    r8 uint64
    rax uint64
    rcx uint64
    rdx uint64
    rsi uint64
    rdi uint64
    orig_rax uint64
    rip uint64
    cs uint64
    eflags uint64
    rsp uint64
    ss uint64
    fs_base uint64
    gs_base uint64
    ds uint64
    es uint64
    fs uint64
    gs uint64
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
	for i := 0; i < len(slice); i++ {
		v := slice[i]
		fmt.Println("i: ", i, "v: ", v)

		if v == "-h" {
			param.helpSet = true
		} else if v == "-p" {
			param.pidSet = true

			if i+1 >= args_length-1 {
				fmt.Println("ERROR couldn't find pid.")
				os.Exit(1)
			}

			pid_str := slice[i+1]
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
	c_pid := C.int(pid)

	fmt.Println("Starting ptrace on pid:", pid)

	ret := C.wrapper_connect(c_pid)

	if ret < 0 {
		fmt.Println("ERROR could not connect to pid: ", pid)
		os.Exit(1)
	}

	i := 0
	for {
		//fmt.Println("loop itr: ", i)
		i++

		c_status := C.int(0)
		C.wrapper_waitpid(c_pid, &c_status)

		if C.wrapper_has_exited(c_status) != 0 {
			fmt.Println("Target process exited")
			break
		}

		sys_int := SysCallType((C.wrapper_get_sys_call(c_pid)))

        if sys_int == write {
            write_sys_str := parseWrite(int(c_pid))
		    fmt.Println(write_sys_str);
        } else {
		    fmt.Println(get_sys_call(sys_int), "(--) = ret")
        }

		C.wrapper_continue(c_pid)
	}

	//ret = C.wrapper_detach(C.int(pid))

	//if ret < 0 {
	//	fmt.Println("ERROR could not detach to pid: ", pid)
	//	os.Exit(1)
	//}
}

func parseWrite(c_pid int) string {
    cstr := C.wrapper_peek_regs(C.int(c_pid))

    gostr := C.GoString(cstr)
    gostr = strings.ReplaceAll(gostr, "\n", "\\n")

    C.free(unsafe.Pointer(cstr))

    return "write(1, " + gostr + ", --) = --"
}
