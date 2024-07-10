package main

import "fmt"

func main() {
	a, b := 1, 100
	for a < b {
		a *= 2
		//fmt.Println(a)
	}
	fmt.Println(a)
}
