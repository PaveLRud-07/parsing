package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func CountingLinks() {
	//записываем в переменную данные запроса гет
	resp, err := http.Get("https://www.packtpub.com/")
	if err != nil {
		panic(err)
	}
	//записываем в переменую данные тела из ответа гет
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// преврашщаем срез байтов в строку
	stringBody := string(data)
	// создаём файл в который записываем срез байт
	var out *os.File
	out, err = os.OpenFile("index.txt", os.O_CREATE|os.O_WRONLY, 0664)
	out.Write(data)
	//считаем сколько ссылок на странице
	numLinks := strings.Count(stringBody, "<a")
	fmt.Printf("PAckt Publishing homepage has %d links!\n", numLinks)
}
