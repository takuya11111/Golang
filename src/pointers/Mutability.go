package main

import "fmt"

type person struct {
	name string
}

/*
func main() {
	p := person{"Richard"}
	rename(p)
	fmt.Println(p) // output: Richard
}

func rename(p person) {
	p.name = "test"
}
*/

func main() {
	p := person{"Richard"}
	rename(&p)
	fmt.Println(p) // output: test
}

func rename(p *person) {
	p.name = "test"
}
