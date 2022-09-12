package main

import (
	"fmt"
	"time"
	"rsc.io/quote"
)

func RandomQuote() {
	fmt.Println(quote.Opt())
}

func hello() {
	fmt.Println("Hello, world.")
}

func timeCantell() {
	fmt.Println(time.Now().UTC().Format(time.RFC1123))
}
func main() {
	hello()
	RandomQuote()
	timeCantell()
}
