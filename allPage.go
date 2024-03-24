package main

import (
	"log"
	"net/http"
	"os"
)

func allPAge() {
	var r *http.Response
	var err error
	r, err = http.Get("https://www.fon.bet/sports/football/")
	if err != nil {
		panic(err)
	}
	if r.StatusCode == 200 {
		//создаём срез байтов который потом запишем в оут
		var webPageContent []byte
		// длина среза
		var bodyLength int = 1270
		//зададим длину байтов и присвоим им значения по умолчанию
		webPageContent = make([]byte, bodyLength)
		//считыве даныне из тела в среайтов вбпейдж контент
		r.Body.Read(webPageContent)
		var out *os.File
		out, err = os.OpenFile("index.txt", os.O_CREATE|os.O_WRONLY, 0664)

		if err != nil {
			panic(err)
		}
		//запишем в файл даные из среза байтов
		out.Write(webPageContent)
		//закрываем файл
		out.Close()
	} else {
		log.Fatal(r.Status)
	}

}
