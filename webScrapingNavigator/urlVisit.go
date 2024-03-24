package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// функция которя отвечает за удаление префикса
func retOnlyUrl(a []string) []string {
	var b []string
	for _, v := range a {
		b = append(b, strings.TrimPrefix(v, "<a href=\"https://hub.packtpub.com/"))
	}
	return b
}

// функция рекурсии для запуска каждый раз когда заходим на страницу
func recursion(a map[string]int, b string) {
	res, err := http.Get("https://hub.packtpub.com/" + b)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	// поиск до первого вхождения благодаря ?
	re := regexp.MustCompile("<a href=\"https://hub.packtpub.com/.+?/\"")
	priceMatches := re.FindAllString(string(data), -1)
	e := retOnlyUrl(priceMatches)
	for _, v := range e {
		if a[v] == 1 {
			continue
		} else {
			a[v] = 1
			recursion(a, v)
		}
	}
}
func main() {
	visitedMap := make(map[string]int)
	res, err := http.Get("https://hub.packtpub.com/")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile("<a href=\"https://hub.packtpub.com/.+?/\"")
	priceMatches := re.FindAllString(string(data), -1)
	//создаёться масив найденных строк на первой странице
	e := retOnlyUrl(priceMatches)
	fmt.Println(e)
	for _, v := range e {
		//если строка имеет префикс  тогда пропучскаем ееё
		if visitedMap[v] == 1 {
			continue
		} else {
			// присваиваем посешёной странице нужный префикс
			visitedMap[v] = 1
			//открываем эту страницу в функции посещения её
			recursion(visitedMap, v)

			if err != nil {
				panic(err)
			}

		}
	}
	fmt.Println(visitedMap, 2131)
}
