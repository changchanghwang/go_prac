package main

import (
	"fmt"
	"time"
)

var a = 1

func main() {
	var b = 2
	c := 3
	const d = 4

	if c == 3 {
		e := 5
		fmt.Println(a, b, c, d, e)
	}

	fmt.Println(a, b, c, d)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//  while loop
	// for  a < 10 {
	// 	fmt.Println(a)
	// }

	// for {
	// 	fmt.Println("Infinite loop")
	// }
	test()

	aaa := A("hello World")
	fmt.Println(aaa)
}

func test() {
	defer fmt.Println("hello world")
	defer fmt.Println("hello world2")

	time.Now()
	defer fmt.Println("hello world3")
}

func A[T any](item T) T {
	return item
}
