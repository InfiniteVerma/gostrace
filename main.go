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

// RDI: First parameter
// RSI: Second parameter
// RDX: Third parameter
// R10: Fourth parameter
// R8: Fifth parameter
// R9: Sixth parameter
type UserRegsStruct struct {
	r15      uint64
	r14      uint64
	r13      uint64
	r12      uint64
	rbp      uint64
	rbx      uint64
	r11      uint64
	r10      uint64
	r9       uint64
	r8       uint64
	rax      uint64
	rcx      uint64
	rdx      uint64
	rsi      uint64
	rdi      uint64
	orig_rax uint64
	rip      uint64
	cs       uint64
	eflags   uint64
	rsp      uint64
	ss       uint64
	fs_base  uint64
	gs_base  uint64
	ds       uint64
	es       uint64
	fs       uint64
	gs       uint64
}

func debugUserRegsStruct(data UserRegsStruct) {
	fmt.Println(
		"r15 : ", data.r15,
		"\nr14 : ", data.r14,
		"\nr13 : ", data.r13,
		"\nr12 : ", data.r12,
		"\nrbp : ", data.rbp,
		"\nrbx : ", data.rbx,
		"\nr11 : ", data.r11,
		"\nr10 : ", data.r10,
		"\nr9 : ", data.r9,
		"\nr8 : ", data.r8,
		"\nrax : ", data.rax,
		"\nrcx : ", data.rcx,
		"\nrdx : ", data.rdx,
		"\nrsi : ", data.rsi,
		"\nrdi : ", data.rdi,
		"\norig_rax : ", data.orig_rax,
		"\nrip : ", data.rip,
		"\ncs : ", data.cs,
		"\neflags : ", data.eflags,
		"\nrsp : ", data.rsp,
		"\nss : ", data.ss,
		"\nfs_base : ", data.fs_base,
		"\ngs_base : ", data.gs_base,
		"\nds : ", data.ds,
		"\nes : ", data.es,
		"\nfs : ", data.fs,
		"\ngs : ", data.gs,
	)
}

type SysCallInfo struct {
	name   string
	param1 string
	param2 string
	param3 string
	ret    string
}

func getSysCallInfoStr(info SysCallInfo) string {
	return info.name + "(" + info.param1 + ", " + info.param2 + ", " + info.param3 + ") = " + info.ret
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

	var sysCallPair [2]UserRegsStruct
	for {
		c_status := C.int(0)
		C.wrapper_waitpid(c_pid, &c_status)

		if C.wrapper_has_exited(c_status) != 0 {
			fmt.Println("Target process exited")
			break
		}

		sysInt := SysCallType((C.wrapper_get_sys_call(c_pid)))

		if sysCallPair[0].orig_rax == 0 {
			//fmt.Println("First sys call")

			sysCallPair[0] = getRegisters(c_pid)
		} else {
			if sysCallPair[0].orig_rax != uint64(sysInt) {
				//fmt.Println("Second sys call doesn't match first. resetting")
				sysCallPair[0] = UserRegsStruct{}
				sysCallPair[1] = UserRegsStruct{}

				sysCallPair[0] = getRegisters(c_pid)
			} else {
				sysCallPair[1] = getRegisters(c_pid)
			}
		}

		//fmt.Println(sysCallPair)

		if sysCallPair[1].orig_rax != 0 {

			printSysCall(c_pid, sysCallPair)
			//fmt.Println("Processed a sys call pair, resetting")
			sysCallPair[0] = UserRegsStruct{}
			sysCallPair[1] = UserRegsStruct{}
		}

		C.wrapper_continue(c_pid)
	}

	//ret = C.wrapper_detach(C.int(pid))

	//if ret < 0 {
	//	fmt.Println("ERROR could not detach to pid: ", pid)
	//	os.Exit(1)
	//}
}

func printSysCall(c_pid C.int, sysCallPair [2]UserRegsStruct) {

	sysCall := SysCallType(sysCallPair[1].orig_rax)
	rsi := C.long(sysCallPair[1].rsi)
	rdx := C.ulong(sysCallPair[1].rdx)

	sysCallInfo := SysCallInfo{
		name:   get_sys_call(SysCallType(sysCallPair[1].orig_rax)),
		param1: fmt.Sprint(sysCallPair[1].rdi),
		param2: "",
		param3: fmt.Sprint(sysCallPair[1].rdx),
		ret:    fmt.Sprint(sysCallPair[1].rax),
	}

	processed := false
	switch sysCall {
	case write:
		c_buffer_ptr := C.wrapper_get_data(c_pid, rsi, rdx)
		if c_buffer_ptr == nil {
			fmt.Println("ERROR c_buffer_ptr is NULL!")
			os.Exit(1)
		}
		defer C.free(unsafe.Pointer(c_buffer_ptr))
		buffer := strings.ReplaceAll(C.GoString(c_buffer_ptr), "\n", "\\n")
		sysCallInfo.param2 = buffer
		fmt.Println(getSysCallInfoStr(sysCallInfo))
		processed = true
	}

	if processed == false {
		fmt.Println(sysCallInfo.name)
	}
}

func getRegisters(c_pid C.int) UserRegsStruct {
	c_struct_ptr := C.wrapper_get_regs(c_pid)

	if c_struct_ptr == nil {
		fmt.Println("ERROR c_struct_ptr is NULL!")
		os.Exit(1)
	}

	defer C.free(unsafe.Pointer(c_struct_ptr))

	regs := (*UserRegsStruct)(unsafe.Pointer(c_struct_ptr))

	return *regs
}
