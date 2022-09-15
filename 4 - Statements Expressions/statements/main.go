package main

import f "fmt"

func main() {
	f.Println("Hello!")
	xFunc()

	// Statements are instructions that perform actions.
	// They are composed of expressions.
	// Run code.
	f.Println("Package scope x address", &x)
	x := 20
	f.Println("main scope x address", &x)

	if x > 5 {
		f.Println("x is greater than 5")
	} else {
		f.Println("x is less than 5")
	}

	semicolons()
}
