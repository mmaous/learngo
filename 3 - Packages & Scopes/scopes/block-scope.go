package main

func nope() { // start of block scope
	const a = 3
	var b int = 4
	// _ in golang is a blank identifier, it is used to ignore the value of a variable
	_ = a + b
} // end of block scope

// this gives a warning because main() is declared in the same package as nope() (i.e. main)
func main() { // start of block scope
	// a and b are not visible here

	// ERROR:
	// fmt.Println(a, b)
	nope()
} // end of block scope
