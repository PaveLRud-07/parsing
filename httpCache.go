package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

func hhtCashe() {
	// создаём место где храним кеш файлы
	storage := diskcache.New("./cahe")
	// записываем кеш файлы в переменую
	cache := httpcache.NewTransport(storage)

	cache.MarkCachedResponses = true
	cachedClient := cache.Client()

	println("Caching : example")
	// записываем данные в переменую из структуры клиента
	resp, err := cachedClient.Get("http://www.example.com/index.html")
	if err != nil {
		panic(err)
	}
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	println("Requesting : example")
	resp, err = cachedClient.Get("http://www.example.com/index.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(*resp)
	_, ok := resp.Header["X-From-Cache"]
	if ok {
		println("Result was pulle from the cache!")
	}
}
