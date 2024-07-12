package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	//'_' does not hold an element.
	for _, val := range array {
		if val%2 == 1 {
			fmt.Println(val)
		}
	}
}
