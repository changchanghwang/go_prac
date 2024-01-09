package main

import (
	"fmt"
)

/**
* i는 3,2,1 순서로 들어가지만 defer는 LIFO(Last In First Out)이기 때문에 1,2,3 순서로 출력된다.
 */
func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

/**
* 익명함수 자체가 마지막에 평가되고 실행되기 때문에 0 0 0 으로 된다. -> 이런 문장은 피하는게 좋다.
 */
func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

/**
* 매개변수가 이미 저장이 되어있는 상태에서 익명함수를 마지막에 평가하고 실행하기때문에 1, 2, 3 순서로 출력된다. -> 가장 바람직하다고 한다. 원하는 변수를 명시적으로 전달하기 때문.
 */
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
