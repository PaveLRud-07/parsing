package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
)

func maifn() {
	doc, err := htmlquery.LoadURL("https://www.packtpub.com/free-learning")
	if err != nil {
		panic(err)
	}
	dealTextNodes := htmlquery.Find(doc, `//div[@class="card__inner"]//div[@class="card__content"]`)
	if err != nil {
		panic(err)
	}
	println("Here is freebook today")
	println("------------------------")

	for _, node := range dealTextNodes {
		text := strings.TrimSpace(node.Data)
		for _, v := range node.Attr {
			fmt.Println(v)
			switch v.Key {
			case "card__title":
				fmt.Println(1)
			}
		}

		if text != "" {
			println(text)
		}
	}
}
