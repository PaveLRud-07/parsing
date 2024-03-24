package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

func postMethod() {
	//содаём ключ значение по которуму будут получать методом пост
	data := url.Values{}
	//передаём ключ значение
	data.Set("s", "Golang")
	//используя метод пост передаём данные на сайт в форму поиска
	response, err := http.PostForm("https://hub.packtpub.com/", data)
	fmt.Print(data)
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	// считываем данные потоком тело(структуры)
	datas, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Body)
	//переводим структуру в строковоу представление
	stringBody := string(datas)
	//создам форму поиска пори помощи регулярных выражений
	re := regexp.MustCompile(".*entry-title td-module-title.*/h3")
	//ищем по регулярным выражениям
	priceMatches := re.FindAllString(stringBody, -1)
	fmt.Println(priceMatches)
	println(response.StatusCode)

	var out *os.File
	//создаём или открываем файл индекс
	out, err = os.OpenFile("index.txt", os.O_CREATE|os.O_WRONLY, 0664)
	//записываем в него данные
	out.Write(datas)
}
