package main

import "fmt"

type Author struct {
	ID     int
	Name   string
	Gender string
}

type Book struct {
	ID    int
	Name  string
	Price int
	Author
}

func main() {
	author1 := Author{
		ID:     1,
		Name:   "Ignacio Daniel",
		Gender: "Male",
	}

	book1 := Book{
		ID:     1,
		Price:  1200,
		Name:   "BookA",
		Author: author1,
	}

	fmt.Println(book1)

	p := &book1
	p.Name = "BookB"

	fmt.Println(book1)
}
