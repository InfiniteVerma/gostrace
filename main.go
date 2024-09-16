package main

import (
    "fmt"
    "os"
)

func main() {
    
    args_length := len(os.Args)

    if args_length != 2 {
        fmt.Println("ERROR invalid number of args passed. Usage: ./goptrace <pid>")
        os.Exit(0)
    }

    pid := os.Args[1]

    fmt.Println("Starting ptrace on pid:", pid)
}
