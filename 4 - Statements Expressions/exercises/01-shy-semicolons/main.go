package main

// ---------------------------------------------------------
// EXERCISE: Shy Semicolons
//
//  1. Try to type your statements by separating them using
//     semicolons
//
//  2. Observe how Go fixes them for you
// Solution: 
//  1. Try to type your statements by separating them using semicolons
//  2. Observe how Go fixes them for you with gofmt
//  3. Run the program with gofmt file-name.go 
//  4. Observe how gofmt fixes the code for you
// ---------------------------------------------------------

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground"); fmt.Println("Hello, Again"); fmt.Println("Hello, Again Again");

}
