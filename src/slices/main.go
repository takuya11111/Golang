package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {

	var a = []int{1, 2, 3}
	//  要素0をスライスaの先頭に追加する。
	a = append([]int{0}, a...)
	// -3, -2, -1の要素をスライスaの先頭に追加
	a = append([]int{-3, -2, -1}, a...)
	fmt.Println(a)

	var s []string
	s = append(s, "Apple")
	s = append(s, "Orange", "Lemon")

	fmt.Println(s)

	// 短い構文でslice(スライス)を定義する
	myNumbers := []int{2, 3, 1, 10}
	// slice(スライス)に要素を追加する
	myNumbers = append(myNumbers, 20)
	fmt.Println(myNumbers)

	//slice(スライス)のslice(スライス)
	mySlice := [][]string{
		[]string{"A", "B", "C"},
		[]string{"D", "E", "F"},
	}

	fmt.Println(mySlice)

	// Slice of struct of User
	user1 := User{
		ID:   1,
		Name: "Corey Mullins",
	}
	user2 := User{
		ID:   2,
		Name: "Terrell Green",
	}
	usersSlice := []User{user1, user2}
	fmt.Println(usersSlice)

	// Slice length and capacity
	sl := []int{5, 6, 9, 2, 1}
	fmt.Println("Length", len(sl))
	fmt.Println("Cap", cap(sl))

	// Using make to declare a slice with length = 0 and capacity is 3
	numbers := make([]int, 0, 3)
	fmt.Printf("cap %v, len %v, pointer %p\n", cap(numbers), len(numbers), numbers)
	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)
		fmt.Printf("cap %v, len %v, pointer %p\n", cap(numbers), len(numbers), numbers)
	}

	//Sub slice
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Original: ", original)
	x := original[0:] // sublice from 0 to end of slice, same with original
	fmt.Println("original[0:]", x)

	y := original[:4] // subslice from 0 to 4
	fmt.Println("original[:4]", y)

	z := original[2:8] // subslice from 2 to 8
	fmt.Println("original[2:8]", z)

	//make
	aa := make([]int, 5) // len(a)=5
	fmt.Println(aa)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	fmt.Println(b)

	var ss []int // 入力なしでスライスを宣言する
	fmt.Println(ss, len(ss), cap(ss))
	if ss == nil {
		fmt.Println("nil!")
	}
	// output `nil!`

	for index, item := range mySlice {
		fmt.Printf("Index %d, item: %v\n", index, item)
	}

}
