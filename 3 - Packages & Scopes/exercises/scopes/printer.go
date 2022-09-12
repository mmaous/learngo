package main

import "fmt"

// printer.go can call this function
//
// this is because, bye function is in the package-scope
// of the main package now.
func hello() {
	fmt.Println("My name is Adolf")
	bye() 
}