# gostrace

A process tracer CLI written in golang

# Plan
 
V1
 - [x] Understand scope
 - [x] Basic cli setup
 - [x] Interact with ptrace.h header file 
 - [x] Add a wrapper.h/.c and use it in go using cgo
 - [x] Basic attach to pid and just print sys call name

V2
 - [x] Print 1 syscall per line like strace
 - [ ] Print write sys call properly

V3
 - [ ] Write more syscalls with proper args and ret val
