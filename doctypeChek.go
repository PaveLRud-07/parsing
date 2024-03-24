package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func doctypeChek() {
	//создаём обьект которые принимает даные сайта в ссылке
	resp, err := http.Get("https://www.packtpub.com/")
	//проверяем ошибку
	if err != nil {
		panic(err)
	}
	// получаем данные из нашего обьекта (тело обьекта )
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//преврашаем данные из срезай байтов в строки
	stringBody := strings.ToLower(string(data))
	// в зависимости от содержания обьекта возвращаем ответ с версией хтмл
	if strings.Contains(stringBody, "<!doctype html>") {
		fmt.Println("This webpage Html5")
	} else if strings.Contains(stringBody, "html/strict.dtd") {
		fmt.Println("html4 strict")
	} else if strings.Contains(stringBody, "html/loose.dtd") {
		fmt.Println("Transitional")
	} else if strings.Contains(stringBody, "html/frameset.dtd") {
		fmt.Println("framest")
	} else {
		fmt.Println("not found")
	}
}
