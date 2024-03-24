package main

import (
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
)

func htmlqueryS() {
	//загружаем сраицу
	doc, err := htmlquery.LoadURL("https://www.packtpub.com/free-learning")
	if err != nil {
		panic(err)
	}
	//дляtmlquery..Find находитвсе вхождения еслиих больше ё
	dealTextNodes := htmlquery.Find(doc, `//div[@class="card__inner"]//text()`)

	println("Here is freebook today")
	println("------------------------")

	for _, node := range dealTextNodes {
		//записываем в текст удаляем пробелы текст храниться в нодедата
		text := strings.TrimSpace(node.Data)
		// записываем все вхождения регулярных вхождений с пустыми блоками дивы спаны и т.д считаются пробелами
		matchTagNames, _ := regexp.Compile("^(div|span|h2|br|ul|li)$")
		// удаляем из текста т.е заменяем дивы ипереносы на пестоту
		text = matchTagNames.ReplaceAllString(text, "")
		//выводим текст если он не равен пустоте
		if text != "" {
			println(text)
		}
	}

}
