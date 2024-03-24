package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func RegFindPrice() {
	//получаем данные с сайта по средствам метода гет стандартной билблиотеки
	resp, err := http.Get("https://www.sela.ru/eshop/kids/boy/?age%5B%5D=5067")
	if err != nil {
		panic(err)
	}
	//считываем даннык потоком из структуры респ под методом боди
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//переводим жанные в строковый вид
	stringBody := string(data)
	// создаём регулярное выражение по которому будем искать данные в теле
	re := regexp.MustCompile(`.*product-card__price-container.*\n.*[0-9]*.*`)
	//находжим все вхлждения используя метод финдстрингсубматч
	priceMatches := re.FindStringSubmatch(stringBody)
	fmt.Print(priceMatches)
}
