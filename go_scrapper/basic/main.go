package main //컴파일을 위한 패키지

import (
	"example/hello/go_scrapper/basic/something"
	"fmt"
	"strings"
)

func multiply(a int, b int /* a, b int로 작성해도 됨*/) int{
	return a * b
}

// go에서 함수는 여러개를 return 할 수 있다.
func lenAndUpper(name string)(int, string){
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string){
	fmt.Println(words)
}

func main(){
	// fmt.Println("hello, world")
	something.SayHello() //함수 export 를 하고 싶다면 대문자로 시작
	something.SayBye()
	const engName string = "arthur"
	name := "hwang" // === var engName string = "hwang" func 안에서만 사용가능 
	name = "chang"
	fmt.Println(name)
	fmt.Println(engName)
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper(("hwang")) // _로 변수명을 하면 컴파일러가 무시함
	fmt.Println(totalLength,upperName)
	repeatMe("hwang","chang","hwan","arthur")
}
