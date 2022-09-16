package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)
	people := [3]string{"hwang","arthur","chang"}
	for _, person := range people{
		go isSexy(person, c)
	} 
	for i:=0; i< len(people); i++{
		fmt.Println("waiting for messages")
		fmt.Println(<-c)
	}
}

func isSexy(person string, channel chan string){
	time.Sleep(time.Second * 1)
	channel <- person + " is sexy"
}