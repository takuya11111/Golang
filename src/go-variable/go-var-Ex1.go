package main

import "fmt"

func main() {
	// implicitly define a variable n with type int and initial value is 10
	n := 10
	fmt.Println(n)

	// implicitly define a variable n,  has same type with n and initial value is n
	k := n
	fmt.Println(k)

	// implicitly define a variable 'firstName, lastName` with type string and initial values are "Ursus", "Lennon"
	firstName, lastName := "Ursus", "Lennon"
	fmt.Println(firstName, lastName)
}
