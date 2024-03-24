package main

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func dailyDeals() {
	// в css необязательно находить точный класс
	//вписываем в переменную doc работу метода newDocument
	// которая приниамает в качестве параметра ссылку на страницу
	// после чего возвращает структуру документа которая хранит в себе
	doc, err := goquery.NewDocument("https://web.archive.org/web/20190519183308/https://www.packtpub.com/latest-releases")
	if err != nil {
		panic(err)
	}
	println("Here are the latest releases!")
	println("-----------------------------")
	// ишем при помощи метода find необходимые теги (после каждого нахождения запускает EACH)
	//после чего на каждом найденом теге мы используем метод Each
	//который получает неа вход I типа int  и структуру селекшон
	//есть возможность прописать точный див используя правила написания которые присутствуют в htmlquery
	doc.Find(`div.landing-page-row div[itemtype$="/Product"]`).
		Each(func(i int, e *goquery.Selection) {
			//тут запушен цикл внутри пока не дойдём до конца структуры
			// создаём переменые где будем хранить данные
			var title string
			var price float64
			// записываем даные в титл испольуя даные из структуры
			title, _ = e.Attr("data-product-title")
			// получаем даные цены из атрибута (атрибут это тег хтмл)
			priceString, _ := e.Attr("data-product-price")
			//переводим цену в число
			price, err = strconv.ParseFloat(priceString, 64)
			if err != nil {
				println("Failed to parse price")
			}
			//выводим результат и ищем совпадления дальше
			fmt.Printf("%s ($%0.2f)\n", title, price)
		})
}
