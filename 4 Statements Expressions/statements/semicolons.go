package main

import f "fmt"

func semicolons() {
	//* Golang compiler Error `syntax error: unexpected f at end of statement`
	// f.Println("line 1") f.Println("line 2") //! ERROR "expected ';', found newline"
	//* Correct way to write the above code
	f.Println("line 1"); f.Println("line 2")

}
