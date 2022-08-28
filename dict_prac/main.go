package main

import (
	"example/hello/dict_prac/mydict"
	"fmt"
)

func main(){
	dictionary := mydict.Dictionary{"first": "first word"}
	// 추가
	word := "second"
	err := dictionary.Add(word, "second Word")
	if err !=nil {
		fmt.Println(err)
	}
	// 수정
	err2 := dictionary.Update(word, "seconde run")
	if err2 !=nil{
		fmt.Println(err2)
	}else{
		fmt.Println(word)
	}
	
	// 삭제
	dictionary.Delete("first")

	// 찾기
	definition, err3 := dictionary.Search("first")
	if err3 !=nil {
		fmt.Println(err3)
	}else{
		fmt.Println(definition)
	}
}