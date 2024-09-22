package main

import "fmt"

type TimeSpec struct {
	Sec  int64
	NSec int64
}

func toString(t TimeSpec) string {
	sSec := fmt.Sprint(t.Sec)
	sNSec := fmt.Sprint(t.NSec)
	return "{tv_sec=" + sSec + ", tv_nsec=" + sNSec + "}"
}
