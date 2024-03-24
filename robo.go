package main

import (
	"fmt"
	"net/http"

	"github.com/temoto/robotstxt"
)

func robo() {
	// получаю файл робот и прочую информацию о сервере
	resp, err := http.Get("https://www.packtpub.com/robots.txt")
	if err != nil {
		panic(err)
	}
	fmt.Print(resp)
	// проверяем конкретно тело робот ткст
	data, err := robotstxt.FromResponse(resp)
	if err != nil {
		panic(err)
	}
	//
	var testUrls []string
	//проверяем есть ли в бане го клиент и к чему у него есть доступ
	grp := data.FindGroup("Go-http-clien/1.1")
	if grp != nil {
		testUrls = []string{
			"/ru/",
			"/all?search=Go",
			"/bundles",

			"/contact/",
			"/search/",
			"/user/password/",
		}
	}
	for _, url := range testUrls {
		print("cheking" + url + "...")

		if grp.Test(url) == true {
			println("ok")
		} else {
			println("x")
		}
	}
	fmt.Print(*grp)

}
