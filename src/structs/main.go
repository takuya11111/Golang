package main

import "fmt"

type User struct {
	ID     int
	Name   string
	Gender string
}
type Category struct {
	ID   string
	Name string
}

type Product struct {
	ID       string
	Name     string
	Price    int
	Category *Category
}

func main() {
	// Initialize an variable with User type
	user1 := User{
		ID:     1,
		Name:   "Emma Grant",
		Gender: "Female",
	}
	fmt.Println(user1)

	// Assign new value a fields
	user1.Name = "Lila Payne"
	fmt.Println("user1 after assigned new name:", user1)

	// Access fields value in struct
	fmt.Println(user1.Name)
	fmt.Println(user1.Gender)

	// Define "empty" struct
	user2 := User{}
	user2.ID = 2
	user2.Name = "Ignacio Daniel"
	user2.Gender = "Male"
	fmt.Println(user2)

	// Pointer to struct
	userPointer := &User{
		ID: 3,
	}
	userPointer.Name = "Cedric Henry"
	userPointer.Gender = "Male"

	// Print userPointer memory address
	fmt.Printf("userPointer mem address: %p\n", userPointer)

	// Print pointer value
	fmt.Println("Pointer data: ", *userPointer)

	// Nested struct
	iPhone := Product{
		ID:    "1",
		Name:  "iPhone",
		Price: 1200,
		Category: &Category{
			ID:   "2",
			Name: "Phone",
		},
	}

	fmt.Println("Product info: ", iPhone)               // memory address of category will be show at Category
	fmt.Println("Product category: ", *iPhone.Category) // use `*` operator to denote the pointer's underlying value.
}
