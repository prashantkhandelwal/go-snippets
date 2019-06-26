/**
*   Program to get the environment/system details.
*   Makes use of the runtime package.
*   GoDoc Reference: https://golang.org/pkg/runtime/
**/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Compiler:", runtime.Compiler, "compiler")
	fmt.Println("Machine Architecture:", runtime.GOARCH, "machine")
	fmt.Println("Go version:", runtime.Version())
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	fmt.Println("Number of Goroutines:", runtime.NumGoroutine())
}
