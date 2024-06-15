package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

/**
* go run mastering_go/parse-time/main.go 12:34 와 같이 실행하면 12:34를 파싱하여 출력한다.
**/
func main() {
	var myTime string
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}
	myTime = os.Args[1]

	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	} else {
		fmt.Println(err)
	}
}
