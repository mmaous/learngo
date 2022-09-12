package main

//! this is a error

// This file cannot see main.go's imported names ("fmt").
// Because the imported names belong to file scope.

func bye() {
	fmt.Println("Bye!")
}