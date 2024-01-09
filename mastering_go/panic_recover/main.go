package main

import (
	"fmt"
)

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()!")
		}
	}()
	fmt.Println("About to call b()")
	b() //여기서 panic이 일어나고 defer문으로 정의한 익명함수가 제어권을 찾아왔기 때문에 main은 정상적으로 종료되었다.
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()!")
	panic("Panic in b()!")
	fmt.Println("Exiting b()!")
}

func main() {
	a()
	fmt.Println("main() ended")
}
