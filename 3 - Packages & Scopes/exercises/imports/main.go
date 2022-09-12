package main

import "fmt"
import f "fmt"
import fm "fmt"

// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

var hello = "hello"
var hey = "hey"
var hi = "hi"


func main() {

	fmt.Println(hello)
	fm.Println(hey)
	f.Println(hi)
	
}