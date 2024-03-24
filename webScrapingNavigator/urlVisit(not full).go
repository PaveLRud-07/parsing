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
		b = append(b, strings.TrimPrefix(v, "<a href=\"https://hub.packtpub.com"))
	}
	return b
}

// функция рекурсии для запуска каждый раз когда заходим на страницу
func recursion(a map[string]int, b string) {
	res, err := http.Get("https://hub.packtpub.com" + b)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile("<a href=\"https://hub.packtpub.com.*/\"")
	priceMatches := re.FindAllString(string(data), -1)
	e := retOnlyUrl(priceMatches)
	for _, v := range e {
		if a[v] == 1 {
			continue
		} else {
			a[v] = 1
			fmt.Println(v)
			recursion(a, v)

		}
	}
}
func urlVisit() {
	visitedMap := make(map[string]int)
	res, err := http.Get("https://hub.packtpub.com/")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile("<a href=\"https://hub.packtpub.com.*/\"")
	priceMatches := re.FindAllString(string(data), -1)
	e := retOnlyUrl(priceMatches)
	for _, v := range e {
		if visitedMap[v] == 1 {
			continue
		} else {
			visitedMap[v] = 1
			recursion(visitedMap, v)

			if err != nil {
				panic(err)
			}

		}
	}
	fmt.Println(visitedMap)
}
