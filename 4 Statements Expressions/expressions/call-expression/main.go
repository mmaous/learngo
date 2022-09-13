package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Call expressions are expressions that call functions.
	// Run code.
	fmt.Println("Go version", runtime.Version())
	fmt.Println("Go OS", runtime.GOOS)
	fmt.Println("Go Arch", runtime.GOARCH)
	fmt.Println("Go NumCPU", runtime.NumCPU())
	go fmt.Println("launching a goroutine")
	fmt.Println("Go NumGoroutine", runtime.NumGoroutine())
	fmt.Println("Go NumCgoCall", runtime.NumCgoCall())
}
