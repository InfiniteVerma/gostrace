# gostrace

A process tracer CLI written in golang. 

The design is rather straightforward. We get sys calls via ptrace functions using C wrappers and on each sys call type found, get it's arguments from memory and print it in human readable format.

## Usage

```
go build .
./gostrace -p <pid>
```

```
go build .
./gostrace <path_to_binary>
```

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
