package main

import f "fmt"

// using all of the operators in go

func main() {
	// Arithmetic Operators
	// + - * / % ++ --
	f.Println("Arithmetic Operators")
	f.Println("5 + 5 =", 5+5)
	f.Println("5 - 5 =", 5-5)
	f.Println("5 * 5 =", 5*5)
	f.Println("5 / 5 =", 5/5)
	f.Println("5 % 5 =", 5%5)

	// Assignment Operators
	// = += -= *= /= %= <<= >>= &= ^= |=
	f.Println("Assignment Operators")
	x := 5
	f.Println("x =", x)
	x += 5
	f.Println("x += 5 =", x)
	x -= 5
	f.Println("x -= 5 =", x)
	x *= 5
	f.Println("x *= 5 =", x)
	x /= 5
	f.Println("x /= 5 =", x)
	x %= 5
	f.Println("x %= 5 =", x)
	x <<= 5
	f.Println("x <<= 5 =", x)
	x >>= 5
	f.Println("x >>= 5 =", x)
	x &= 5
	f.Println("x &= 5 =", x)
	x ^= 5
	f.Println("x ^= 5 =", x)
	x |= 5
	f.Println("x |= 5 =", x)

	// Comparison Operators
	// == != > < >= <=
	f.Println("Comparison Operators")
	f.Println("5 == 5 =", 5 == 5)
	f.Println("5 != 5 =", 5 != 5)
	f.Println("5 > 5 =", 5 > 5)
	f.Println("5 < 5 =", 5 < 5)
	f.Println("5 >= 5 =", 5 >= 5)
	f.Println("5 <= 5 =", 5 <= 5)

	// Logical Operators
	// && || !
	f.Println("Logical Operators")
	f.Println("true && true =", true && true)
	f.Println("true && false =", true && false)
	f.Println("true || true =", true || true)
	f.Println("true || false =", true || false)
	f.Println("!true =", !true)

	// Bitwise Operators
	// & | ^ &^ << >>
	f.Println("Bitwise Operators")
	f.Println("5 & 5 =", 5 & 5)
	f.Println("5 | 5 =", 5 | 5)
	f.Println("5 ^ 5 =", 5 ^ 5)

	

	// & - address of x
	f.Println("Address of x =", &x)
	
}

