package main

import "fmt"

func main() {
	// Prime numbers from 0 to 100
	for i := 0; i < 100; i++ {
		counter := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				counter++
			}
		}
		if counter == 1 {
			fmt.Println(i)
		}
	}
}
