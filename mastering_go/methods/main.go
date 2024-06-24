package main

import "fmt"

type twoInts struct {
	X int64
	Y int64
}

func regularFunction(a, b twoInts) twoInts {
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a twoInts) method(b twoInts) twoInts {
	return twoInts{X: a.X + b.X, Y: a.Y + b.Y}
}

func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: 3, Y: -2}

	fmt.Println(regularFunction(i, j))
	fmt.Println(i.method(j))
}
