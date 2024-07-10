package main

import "fmt"

func main() {
	// if only
	a, b := 10, 10
	if a == b {
		fmt.Println("a and b are same")
	}

	// if and else
	m, n := 10, 20
	if m == n {
		fmt.Println("m and n are same")
	} else {
		fmt.Println("m and n are not same")
	}

	// if , else if, and else
	x, y := 1, 2
	if x > y {
		fmt.Println("x is greater than y")
	} else if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is equal to b")
	}

	totalVNMembers, totalJPMembers := 20, 30
	// If short statement
	if total := totalVNMembers + totalJPMembers; total > 50 {
		fmt.Printf("Total member is %d, to many members\n", total)
	} else if total >= 30 {
		fmt.Printf("Total member is %d, enough\n", total)
	} else {
		fmt.Printf("Total member is %d, not enough\n", total)
	}

}
