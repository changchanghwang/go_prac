package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value int64 = 5
	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2) // -930866580 -> 5434123412312431212의 32bit값. 32bit pointer로는 64bit값을 다 표현할 수 없다.
	*p1 = 54341234
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
