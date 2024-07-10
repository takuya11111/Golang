package main

import "fmt"

func main() {
	var url string = "Hello, world!"
	var totalPhones int

	url = "https://golang.org"
	totalPhones = 10

	// Define a list of variables.
	var totalBooks, totalLaptops int = 1, 2

	// Assign multi values to multi variables
	totalBooks, totalLaptops = 5, 12

	fmt.Println(totalBooks, totalLaptops, totalPhones, url)
}
