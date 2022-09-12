package main

import "rsc.io/quote"
import "fmt" 

func RandomQuote() {
	fmt.Println(quote.Opt())
}

func hello() {
	fmt.Println("Hello, world.")
}
func main() {
	hello()
	RandomQuote()
}
