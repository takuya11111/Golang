package main

import "fmt"

func main() {
	a, b := 2, 1

	fmt.Println("a + b = ", a+b)
	fmt.Println("a - b = ", a-b)
	fmt.Println("a * b = ", a*b)
	fmt.Println("a / b = ", a/b)
	fmt.Println("a % b = ", a%b)
	a++
	fmt.Println("a++ = ", a)
	b--
	fmt.Println("b-- = ", b)
}

// output:
// a + b =  3
// a - b =  1
// a * b =  2
// a / b =  2
// a % b =  0
// a++ =  3
// b-- =  0
