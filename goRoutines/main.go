package main

import (
	"fmt"
	"time"
)

func main(){
	go Count("hwang")
	go Count("arthur") // 해당부분에서 Go routines를 진행하면 메인함수가 끝나버림. 백그라운드에서 실행하는것과 같다고 생가갛면 좋을 것 같음.
	time.Sleep(time.Second * 5)
}

func Count(person string){
	for i:=0;i<10; i++{
		fmt.Println(person,"is sexy", i)
		time.Sleep(time.Second)
	}
}