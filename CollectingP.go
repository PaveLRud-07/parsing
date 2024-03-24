package main

import (
	"bufio"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func CollectingP() {
	//записываем возврат функции goquery.NewDocument("html.adress") возвращает структуру
	doc, err := goquery.NewDocument("https://www.packtpub.com/free-learning")
	if err != nil {
		panic(err)
	}

	println("Here is the free book of the day!")
	println("----------------------------------")
	//в переменую rawText записываем данные Работы метода Find()
	//который получает на вход класс после чего получает на вход уточняющий класс
	//из которого будут браться данные в виде текста
	// после чего используя конструкцию типа div:not(.class) удаляем из выдачи определёный возврат
	rawText := doc.Find(`div.product__info div.product-info__content div:not(.product-info__rating)`).Text()
	// записываем срез байт
	reader := bufio.NewReader(strings.NewReader(rawText))
	var line []byte
	for err == nil {
		line, _, err = reader.ReadLine()
		trimmedLine := strings.TrimSpace(string(line))
		if trimmedLine != "" {
			println(trimmedLine)
		}
	}
}
