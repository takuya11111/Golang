package main

import "fmt"

func main() {
	persons := []string{
		"Herman Matthews",
		"Corey Mullins",
		"Dwight Ramsey",
		"Zachary Carr",
		"Howard Oliver",
		"Brandon Fields",
		"Andy Thornton",
		"Jeff Peters",
		"Terrell Green",
		"Jeffrey Harvey",
	}

	for _, val := range persons {
		if val == "Zachary Carr" {
			fmt.Println(val)
		}
	}

	subPersons := persons[4:8]
	fmt.Println(subPersons)
}
