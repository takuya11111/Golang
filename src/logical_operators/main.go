package main

import "fmt"

func main() {
	fmt.Println("true && true gives ", true && true)
	fmt.Println("true && false gives ", true && false)
	fmt.Println("false && true gives ", false && true)
	fmt.Println("false && false gives ", false && false)
	fmt.Println("true || true gives ", true || true)
	fmt.Println("true || false gives ", true || false)
	fmt.Println("false || true gives ", false || true)
	fmt.Println("false || false gives ", false || false)
	fmt.Println("!false gives ", !false)
	fmt.Println("!true gives ", !true)
}
