package main

import "fmt"

func main() {
	var i int = 10
	// "p" points to i
	p := &i

	fmt.Println("Memory address of p", p)
	fmt.Println("Read i though pointer p", *p)

	// set i through the pointer p
	*p = 20
	fmt.Println(i) // output: 20
}
