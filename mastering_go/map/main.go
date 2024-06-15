package main

import "fmt"

/**
* go의 map은 해시 테이블과 같다. 모든 데이터 타입을 키로 사용할 수 있지만 키에 지정할 데이터 타입은 반드시 비교할 수 있어야 한다.
* (==연산에 적용 가능해야한다. 부동 소수점 수를 키 값으로 사용하면 프로세서나 OS에 의해 정밀도가 차이나서 문제가 발생할 수 잇다.)
* go 언어의 맵은 해시 테이블에 대한 참조값이다. -> 복잡도를 낮출 수 있다.
**/
func main() {
	iMap := make(map[string]int) //string 타입을 키로 사용하는 int 타입의 값을 가진 맵을 생성한다.
	iMap["key1"] = 12
	iMap["key2"] = 23
	fmt.Println("map:", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	} // map literal을 이용하여 들어갈 맵을 직접 지정할 수 잇다.
	fmt.Println("assignedMap:", anotherMap)

	delete(anotherMap, "k1") // map에서 키를 삭제할 때는 delete() 함수를 사용한다.
	fmt.Println("deletedMap:", anotherMap)
	fmt.Println("----------")

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("deletedMap:", anotherMap) // delete 함수는 여러번 호출해도 에러가 발생하지 않는다.

	_, ok := iMap["doesItExist"]
	if ok {
		fmt.Println("Exist!")
	} else {
		fmt.Println("Doesn't exist!")
	}

	for key, value := range iMap {
		fmt.Println(key, value)
	}
}
