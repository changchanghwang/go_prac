package main

import (
	"fmt"
	"unsafe"
)

/*
*
* 배열의 각 원소를 포인터를 이용해 접근하는 방법.
unsafe를 사용했기때문에 하나 더 출력해도 문제가 없고 이상한 값을 리턴한다.
*/
func main() {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]              // array[0]에 대한 메모리주소
	fmt.Println("pointer:", *pointer) // pointer를 역참조 -> 그 주소에 저장되어 있던 값을 리턴한다.
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])

	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Println("pointer:", *pointer)
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	}

	fmt.Println()
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Println("one more:", *pointer)
	memoryAddress = uintptr(unsafe.Pointer(pointer)) + unsafe.Sizeof(array[0])
	fmt.Println("memoryAddress:", memoryAddress)

}
