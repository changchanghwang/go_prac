package main

import (
	"fmt"
)

func reSlice() {
	s1 := make([]int, 5)
	reSlice := s1[1:3] // s1[1] ~ s1[2] 의 주소값을 참조하는 슬라이스를 생성한다.
	fmt.Println(s1)
	fmt.Println(reSlice)

	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice) // s1과 reSlice는 같은 주소값을 참조하고 있기 때문에 값이 변경된다.
}

func printSlice(x []int) {
	for _, number := range x {
		fmt.Print(number, " ")
	}
	fmt.Println()
}

func main() {
	// reSlice()
	aSlice := []int{-1, 0, 4}
	fmt.Printf("aSlice: ")
	printSlice(aSlice)

	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -2)
	aSlice = append(aSlice, -3)
	aSlice = append(aSlice, -4)
	fmt.Printf("aSlice: ")
	printSlice(aSlice)
	fmt.Printf("Cap: %d, Length: %d\n", cap(aSlice), len(aSlice))
}
