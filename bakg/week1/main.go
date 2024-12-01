package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello world!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8989")
	if err := http.ListenAndServe(":8989", nil); err != nil {
		// 에러를 stdout으로 출력하고 os.exit(1)을 호출하여 프로그램을 종료한다.
		// panic은 recover로 복구가 가능하지만 log.Fatal은 복구가 불가능하다.
		// 프로그램 종료되기 전에 반드시 실행되어야 되는 작업이 있다면 log.Fatal을 사용하면 안된다.
		// graceful shutdown? https://emretanriverdi.medium.com/graceful-shutdown-in-go-c106fe1a99d9
		log.Fatal(err)
	}
}
