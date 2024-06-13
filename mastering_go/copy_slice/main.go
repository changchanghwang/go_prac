package main

import "fmt"

func main() {
	a6 := []int{-10, 1, 2, 3, 4, 5}
	a4 := []int{-1, -2, -3, -4}
	fmt.Println("a6:", a6) // a6: [-10 1 2 3 4 5]
	fmt.Println("a4:", a4) // a4: [-1 -2 -3 -4]

	copy(a6, a4)
	fmt.Println("a6:", a6) // a6: [-1 -2 -3 -4 4 5] 앞에 4개 원소만 복사되고 뒤에 2개는 남아서 그대로 출력된다.
	fmt.Println("a4:", a4) // a4: [-1 -2 -3 -4]
	fmt.Println()

	array4 := [4]int{4, -4, 4, -4}
	s6 := []int{1, 1, -1, -1, 5, -5}
	copy(s6, array4[0:])
	fmt.Println("array4:", array4[0:]) // [4 -4 4 -4]
	fmt.Println("s6", s6)              // [4 -4 4 -4 5 -5]
	fmt.Println()

	array5 := [5]int{5, -5, 5, -5, 5}
	s7 := []int{7, 7, -7, -7, 7, -7, 7}
	copy(array5[0:], s7) //copy는 인자로 slice만 받기 때문에 array -> slice로 변환해서 사용해야 한다.

	fmt.Println("array5:", array5) // [7,7,-7,-7,7]
	fmt.Println("s7:", s7)         // [7,7,-7,-7,7,-7,7]
	fmt.Println()
}
