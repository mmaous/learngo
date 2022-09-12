package main // start of package scope

import "fmt"

var declareMeAgain = "I'm declared for the first time"

func anotherScope() {

	declareMeAgain := "I'm declared Again!!!!"

	fmt.Println(declareMeAgain)
}

func main() {

	fmt.Println(declareMeAgain)

	// package-level declareMeAgain isn't effected
	// from the change inside the nested func
	anotherScope()

	fmt.Println(declareMeAgain)

}
