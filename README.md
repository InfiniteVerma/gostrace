# gostrace

A process tracer CLI written in golang. 

The design is rather straightforward. We get sys calls via ptrace functions using C wrappers and on each sys call type found, get it's arguments from memory and print it in human readable format.

## Demo



https://github.com/user-attachments/assets/059e0bb3-cf1f-419b-8aaa-fd3862c30699



## Setup

```
echo 0 | sudo tee /proc/sys/kernel/yama/ptrace_scope
```

## Sys calls implemented:
 - write
 - clock_nanosleep
 - close
 - lseek

## Plan
 
V1
 - [x] Understand scope
 - [x] Basic cli setup
 - [x] Interact with ptrace.h header file 
 - [x] Add a wrapper.h/.c and use it in go using cgo
 - [x] Basic attach to pid and just print sys call name

V2
 - [x] Print 1 syscall per line like strace
 - [x] Print write sys call properly

V3
 - [ ] Write more syscalls with proper args and ret val
