package main

import (
	"fmt"
	"os"
)

func get_sys_call(sys_int SysCallType) string {

	i, ok := sysCallNames[sys_int]

	if ok != true {
		fmt.Println("ERROR could not parse sys_int: ", sys_int)
		os.Exit(1)
	}
	return i
}

func searchStr(slice []string, str string) (bool, int) {
	for i, v := range slice {
		if v == str {
			return true, i
		}
	}
	return false, -1
}
