package main

import "fmt"

func main() {

	totalVNMembers, totalJPMembers := 20, 30
	totalMembers := totalJPMembers + totalVNMembers

	switch totalMembers {
	case 50:
		fmt.Println("Your teams will play at area A")
	case 30:
		fmt.Println("Your teams will play at area B")
	case 20:
		fmt.Println("Your teams will play at area C")
	default:
		fmt.Println("Your team members must be 50, or 30, or 20")
	}
}
