package main

import "fmt"

func main() {
	// sum of odd numbers
	sumOfOdd := 0
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}
		sumOfOdd++
	}
	fmt.Println(sumOfOdd)

	// Output number from 1 to 10.
	min, max := 0, 10
	for {
		fmt.Println("Execute me until min is qual to max")
		min++
		if min == max {
			fmt.Println("Min is qual to max")
			break
		}
	}
}
