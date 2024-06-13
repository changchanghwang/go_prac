package main

import "fmt"

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("aSlice:", aSlice)

	integer := make([]int, 2)
	fmt.Println("integer:", integer)

	integer = nil
	fmt.Println("integer:", integer)
	fmt.Println("--------------------------")

	anArray := [5]int{-1, -2, -3, -4, -5}
	refAnArray := anArray[:] // 기존 배열을 참조하는 새로운 슬라이드 생성

	fmt.Println("anArray:", anArray)
	fmt.Println("refAnArray:", refAnArray)

	anArray[4] = -100
	fmt.Println("anArray:", anArray)
	fmt.Println("refAnArray:", refAnArray) // anArray와 refAnArray는 같은 주소값을 참조하고 있기 때문에 값이 변경된다.
	fmt.Println("--------------------------")

	s := make([]byte, 5)       // slice는 go에서 자동으로 초기화해줌. 정수형 값은 9으로, slice에 대해서는 nil로 초기화한다.
	fmt.Println("s:", s)       // [0 0 0 0 0]
	twoD := make([][]int, 3)   // 다차원 slice의 원소는 슬라이스이기 때문에 nil로 초기화된다.
	fmt.Println("twoD:", twoD) // [[] [] []]
	fmt.Println("--------------------------")

	for i := 0; i < len(twoD); i++ {
		for j := 0; j < 2; j++ {
			twoD[i] = append(twoD[i], i*j) // append() 함수를 사용하여 확장시켜야한다. 존재하지 않는 인덱스를 참조하면 런타임 에러가 발생한다.
		}
	}

	for _, x := range twoD {
		for i, y := range x {
			fmt.Println("i:", i, "value:", y)
		}
		fmt.Println()
	}
}
