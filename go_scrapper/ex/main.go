package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?&searchword=typescript&recruitSort=relation&recruitPageCount=40&inner_com_type=&company_cd=0%2C1%2C2%2C3%2C4%2C5%2C6%2C7%2C9%2C10&show_applied=&quick_apply=&except_read=&ai_head_hunting=&mainSearch=n"

func main() {
	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}
}

func getPage(pageNumber int) {
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(pageNumber+1)
	fmt.Println("Requesting:", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkStatusCode(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	seachCards := doc.Find(".item_recruit")
	seachCards.Each(func(i int, card *goquery.Selection) {
		fmt.Println(card.Find("div").Html())
	})
}

func getPages() int {
	pages := 0

	res, err := http.Get(baseURL)
	checkErr(err)
	checkStatusCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatusCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed with Status:", res)
	}
}
