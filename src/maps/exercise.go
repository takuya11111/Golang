package main

import (
	"fmt"
)

func main() {
	input := "abcdaafd"
	output := map[string]int{}

	for _, char := range input {

		if _, ok := output[string(char)]; ok {
			output[string(char)] = output[string(char)] + 1
			//fmt.Println(output)
		} else {
			output[string(char)] = 1
			//fmt.Println(output)
		}

	}

	fmt.Println(output)
}
