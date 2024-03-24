package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func dailyDeals() {
	doc, err := goquery.NewDocument("https://www.packtpub.com/free-learning")
	if err != nil {
		panic(err)
	}
	println("Here are the latest releases!")
	println("-----------------------------")
	time.Sleep(1 * time.Second)
	rav := doc.Find(`div.grid div.card`)
	rav.Each(func(i int, e *goquery.Selection) {
		var title, description, author string
		link, _ := e.Find(`a`).Attr("href")
		link = "https://www.packtpub.com" + link
		bookInfo, err := goquery.NewDocument(link)
		if err != nil {
			panic(err)
		}
		title = strings.TrimSpace(bookInfo.Find(`div.product-info__content h1.product-info__title`).Text())
		authorNodes := bookInfo.Find(`span.free_learning__author`)
		if len(authorNodes.Nodes) < 1 {
			return
		}
		author = strings.TrimSpace(authorNodes.Nodes[0].FirstChild.Data)
		author = strings.TrimSpace(stringInString(author))
		description = strings.Trim(strings.TrimSpace(e.Find(`p`).Text()), "By")
		fmt.Printf("%s\n%s\n%s\n", title, author, description)
		time.Sleep(1 * time.Second)

	})
}
func stringInString(a string) string {
	e := strings.Split(a, ",")
	var s string
	for _, v := range e {
		s = s + strings.Trim(v, "By \n") + " "
	}
	return s
}
