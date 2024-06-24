package main

import "fmt"

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

type first struct{}

func (a first) F() {
	a.shared()
}
func (a first) shared() {
	fmt.Println("This is shared() from first!")
}

type second struct {
	first
}

func (a second) shared() {
	fmt.Println("This is shared() from second!")
}

func main() {
	var i c
	i.A.A()
	i.B.A()

	first{}.F()
	second{}.shared()
	ii := second{}
	(ii.first).F()
}
