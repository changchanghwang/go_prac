package main

import "fmt"

type myStructure struct {
	Name    string
	Surname string
	height  int32
}

func createStruct(name, surName string, height int32) *myStructure {
	if height > 300 {
		height = 0
	}
	return &myStructure{
		Name:    name,
		Surname: surName,
		height:  height,
	}
}

func retStructure(name, surName string, height int32) myStructure {
	if height > 300 {
		height = 0
	}
	return myStructure{
		Name:    name,
		Surname: surName,
		height:  height,
	}
}

func main() {
	s1 := createStruct("Mihalis", "Tsoukalos", 123)
	s2 := retStructure("Mihalis", "Tsoukalos", 123)
	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

	pS := new(myStructure) // make와 다른점은 make로 생성한 변수는 정상적으로 초기화되고 new로 생성한 변수는 할당된 메모리 공간에 단지 0만 채운다. 그리고 new는 포인터를 리턴한다.
	fmt.Println(pS)
}
