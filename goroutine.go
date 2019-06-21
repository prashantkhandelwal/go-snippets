package main

import "fmt"
import "time"

func f(from string) {
	for index := 0; index < 10; index++ {
		time.Sleep(1000)
		fmt.Println(from, ":", index)
	}
}

func main() {

	go f("Khandelwal")
	f("Prashant")

	go func(msg string){
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
}