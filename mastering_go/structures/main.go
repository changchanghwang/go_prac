package main

import "fmt"

func main() {

	type XYZ struct {
		X int
		Y int
		Z int
	}

	var s1 XYZ
	fmt.Println(s1.Y, s1.Z)

	p1 := XYZ{23, 12, -2}
	p2 := XYZ{Z: 12, Y: 13}
	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := [4]XYZ{}
	pSlice[2] = p1
	pSlice[0] = p2

	fmt.Println(pSlice)

	p2 = XYZ{1, 2, 3}
	fmt.Println(pSlice) // 구조체 배열에 구조체를 할당하면 그 내용이 배열에 복사되는 것이기 때문에 p2의 값이 바뀌어도 pSlice에는 영향을 미치지 않는다.
}
