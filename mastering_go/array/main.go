package main

import "fmt"

func forLoop() {
	for i := 0; i < 10; i++ {
		println(i)
	}
}

func whileLoop1() {
	for {
		println("Infinite loop")
	}
}

func whileLoop2() {
	for ok := true; ok; ok = true { // ok = 표현식으로 while(표현식) 과 같은 효과를 낸다.
		println("Infinite loop")
	}
}

func Array2() {
	anArray := [4]int{0, 1, -1, 2}
	twoDepth := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	threeDepth := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}

	fmt.Println("The length of", anArray, "is", len(anArray))
	fmt.Println("The first of ", twoDepth, "is", twoDepth[0][0])
	fmt.Println("The first of ", threeDepth, "is", threeDepth[0][0][0])
}

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}
	fmt.Println()
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println()
	j := 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if j > 10 {
			anExpression = false
		}
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value)
	}

	Array2()
}
