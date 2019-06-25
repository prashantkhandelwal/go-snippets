/**
 *  Program to handle the arguments.
 *  Displays the total number of arguments passed to the program and
 *  display what are the arguments by looping over them.
 *  GoDoc Reference: https://golang.org/pkg/os
**/

package main

import (
	"fmt"
	"os"
)

func main() {
	//By default the program executable is the first argument
	if len(os.Args) == 1 {
		panic("Not enough arguments!")
	} else {
		fmt.Printf("You have passed %d argument(s)", len(os.Args))

		//Replace the '_' with 'i' or any other variable to get the index
		for _, n := range os.Args {
			fmt.Println("")
			fmt.Printf("%v", n)
		}
	}
}
