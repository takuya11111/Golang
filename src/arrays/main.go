package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	// Define array a with 10 element with int type
	var a [3]int
	fmt.Println("Initialize array: ", a)
	// Add element to array a
	a[0] = 3
	a[1] = 10
	a[2] = 5

	fmt.Println(a)

	// we also can use short syntax
	fruits := [4]string{"Apple", "Orange", "Lemon", "Banana"}
	fmt.Println(fruits)

	// Array can be a struct
	usersList := [2]User{}
	usersList[0] = User{ID: 1, Name: "Horace Bennett"}
	usersList[1] = User{ID: 2, Name: "Santiago Mitchell"}
	fmt.Println(usersList)

	for i, user := range usersList {
		// Access i and user here
		fmt.Printf("Index: %d, user: %v\n", i, user)
	}
}
