package main

import "fmt"

func main() {
	// Normal syntax
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["first"] = "My first pair"
	myMap["second"] = "My second pair"
	myMap["third"] = "My third pair"
	fmt.Println(myMap)
	fmt.Println(myMap["third"]) //Access to an element in map

	// Short syntax
	numbersMap := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	numbersMap["4"] = 4
	fmt.Println(numbersMap)

	countriesMap := map[string]string{}
	countriesMap["en-US"] = "US"
	countriesMap["ja-JP"] = "Japan"
	countriesMap["vi-VN"] = "Vietnam"
	fmt.Println(countriesMap)

	// Loop through country map
	for key, value := range countriesMap {
		fmt.Printf("Map key %v, map value: %v\n", key, value)
	}

	if value, ok := countriesMap["en-US"]; ok {
		fmt.Println("Found key", value)
	} else {
		fmt.Println("Not found")
	}

}
