package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan string)
	people := [2]string{"hwang","arthur"}
	for _, person := range people{
		go isSexy(person, c)
	} 
	fmt.Println("waiting for messages")
	fmt.Println(<- c)// channel로 부터 받은 메세지, 블로킹
	fmt.Println(<- c)// channel로 부터 받은 메세지, 블로킹
}

func isSexy(person string, channel chan string){
	time.Sleep(time.Second * 1)
	channel <- person + " is sexy"
}