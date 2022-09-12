package main

import "fmt"

const hello = "Hello, World!"

func main() { // start of block scope
		// a and b are not visible here

	// ERROR:
	// fmt.Println(a, b)
	nope()

	fmt.Println(hello)
} // end of block scope