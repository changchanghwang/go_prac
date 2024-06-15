package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC1123))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	loc2, _ := time.LoadLocation("Asia/Seoul")
	londonTime := t.In(loc)
	seoulTime := t.In(loc2)
	fmt.Println("Paris:", londonTime)
	fmt.Println("Seoul:", seoulTime)
}
