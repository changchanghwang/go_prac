package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int
type t2 int

type MyStruct struct {
	X    int
	Y    float64
	Text string
}

func (myStruct MyStruct) compareStruct(other MyStruct) bool {
	r1 := reflect.ValueOf(&myStruct).Elem()
	r2 := reflect.ValueOf(&other).Elem()

	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func printMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()
	fmt.Println("Type to examine:", t)

	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, " -> ", m)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(200)
	fmt.Println("The type of x1 is ", reflect.TypeOf(x1))
	fmt.Println("The type of x2 is ", reflect.TypeOf(x2))

	var p struct{}
	r := reflect.New(reflect.ValueOf(&p).Type()).Elem()
	fmt.Println("The type of r is ", reflect.TypeOf(r))

	a1 := MyStruct{100, 200.12, "Struct A1"}
	a2 := MyStruct{100, -2, "Struct A2"}

	if a1.compareStruct((a1)) {
		fmt.Println("Equal!")
	}
	if !a1.compareStruct(a2) {
		fmt.Println("Not Equal!")
	}

	var f *os.File
	printMethods(f)
}
