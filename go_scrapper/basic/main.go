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
func lenAndUpper(name string)(length int, uppercase string){
	defer fmt.Println("I'm done") //함수가 끝났을때 defer가 실행됨.
	length = len(name)
	uppercase = strings.ToUpper(name)
	fmt.Println("return start")
	return //이미 변수를 선언해서 반환한다고 해놨으니까 return만 작성해도됨. naked return
}

// ...을 쓰면 args를 한번에 받음
func repeatMe(words ...string){
	fmt.Println(words)
}

func superAdd(numbers ...int) int{
	total := 0
	// for i:=0; i <len(numbers); i++ {
	//   total += numbers[i]
	// } //가능하나 아래처럼도 작성 가능
	for _, number := range numbers{
		total += number
	}
	return total
}

func canIDrink(age int) bool{
	if koreanAge := age + 2 ; koreanAge < 18 { //변수를 만들고 if문 실행 가능 variable expression
		return false
	}
	return true
}

func ageCheck(age int) bool{
	switch koreanAge := age + 2; koreanAge{
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func pointer(){
	a := 2
	b := &a // &은 메모리값을 볼 수 있음
	a = 10
	*b = 20
	fmt.Println(a,b,*b) // a는 값, b는 메모리 주소값 *b 는 해당 메모리의 값
}

func array(){
	arr := [5]int{1,2,3,4,5} //array
	names := []string{"hwang","chang","hwan"} //slice
	names = append(names,"hi")
	fmt.Println(arr,names)
}

func main(){
	// fmt.Println("hello, world")
	something.SayHello() //함수 export 를 하고 싶다면 대문자로 시작
	something.SayBye()
	// const engName string = "arthur"
	// name := "hwang" // === var engName string = "hwang" func 안에서만 사용가능 
	// name = "chang"
	// fmt.Println(name)
	// fmt.Println(engName)
	fmt.Println(multiply(2,2))
	totalLength, upperName := lenAndUpper(("hwang")) // _로 변수명을 하면 컴파일러가 무시함
	fmt.Println(totalLength,upperName)
	repeatMe("hwang","chang","hwan","arthur")

	result := superAdd(1,2,3,4,5,6,7)
	fmt.Println(result)
	
	fmt.Println(canIDrink(16))
	fmt.Println(ageCheck(16))
	
	pointer()
	array()
}
