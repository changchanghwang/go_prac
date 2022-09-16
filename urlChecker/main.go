package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

type urlResult struct {
	url string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan urlResult)
	urls := []string{"https://www.airbnb.com","https://www.naver.com","https://www.google.com","https://www.amazon.com","https://www.reddit.com","https://www.facebook.com"}
	for _, url := range urls{
		go hitURL(url,c )
	}
	for i:=0; i<len(urls); i++{
		result := <-c
		results[result.url] = result.status
	}
	for url,status := range results {
		fmt.Println(url,status)
	}
}

func hitURL(url string, c chan<- urlResult) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400{
		c <- urlResult{url:url, status: "Failed"}
	}
	c <- urlResult{url:url, status: res.Status}
}