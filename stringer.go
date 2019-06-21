/**
*   Program making use of the Stringer interface.
*   Stringer is implemented by any value that has a string method.
*   If you are coming from .NET/JAVA background, 
*   You can think of this method as a overriding the ToString() method.
*   GoDoc Reference: https://golang.org/pkg/fmt/#Stringer
**/

package main

import (
	"fmt"
)

type Person struct {
	Name	string
	Age		uint
}

func (p Person) String() string {
	return fmt.Sprintf("%v is %d years old", p.Name, p.Age)
}

func main() {
	p := Person{
		Name: "Prashant",
		Age: 32,
	}

	fmt.Println(p)
}
