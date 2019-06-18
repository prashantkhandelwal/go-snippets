/**
 *  Code to get the maximum number of parallel operations can be done on a machine.
 *  You can just make use of the runtime.NumCPU() method to get the total number of cores.
 *  https://stackoverflow.com/questions/13234749/golang-how-to-verify-number-of-processors-on-which-a-go-program-is-running
 *  GoDoc Reference: https://golang.org/pkg/runtime/#NumCPU 
**/

package main

import (
	"runtime"
    "fmt"
)

func MaxParallelism() int {
    maxProcs := runtime.GOMAXPROCS(0)
    numCPU := runtime.NumCPU()
    if maxProcs < numCPU {
        return maxProcs
    }
    return numCPU
}

func main() {
	// cores := MaxParallelism()
    // fmt.Printf("%d Cores Available\n", cores)
    fmt.Printf("%d Cores Available\n", runtime.NumCPU())
    fmt.Scanln()
}
