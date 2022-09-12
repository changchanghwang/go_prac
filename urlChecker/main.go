package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	urls := []string{"https://www.airbnb.com","https://www.naver.com","https://www.google.com","https://www.amazon.com","https://www.reddit.com","https://www.facebook.com"}
	for _, url := range urls{
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "Failed"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error{
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
	return errRequestFailed
	}
	return nil
}