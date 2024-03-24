package main

import (
	"fmt"
	"net/http"
	"time"
)

func fiveSec() {
	// создаём карту в которой будем хранить даные с сайта и время обновления
	var lastRequestTime map[string]time.Time = map[string]time.Time{
		"example.com":  time.Time{},
		"packtpub.com": time.Time{},
	}
	// переменая с кол-вом запросов
	maximumNumberOfRequests := 5
	//переменная которая вычисляет паузу
	pageDelay := 5 * time.Second
	// запускаем цикл обновляения
	for i := 0; i < maximumNumberOfRequests; i++ {
		// в зависимости от чётности нечётности цикла запускаем обновление сайта
		if i%2 == 0 {
			// выводим сколько секунд назад обновлён сайт если меньше 5 секунд заслипаем програму
			elapsedTime := time.Now().Sub(lastRequestTime["packtpub.com"])
			fmt.Printf("Elapsed Time: %.2f (s)\n", elapsedTime.Seconds())
			if elapsedTime < pageDelay {
				var timeDiff time.Duration = pageDelay - elapsedTime
				fmt.Printf("Sleeping for %.2f (s)\n", timeDiff.Seconds())
				time.Sleep(timeDiff)

			}
		} else {
			elapsedTime := time.Now().Sub(lastRequestTime["example.com"])
			fmt.Printf("Elapsed Time: %.2f (s)\n", elapsedTime.Seconds())
			if elapsedTime < pageDelay {
				var timeDiff time.Duration = pageDelay - elapsedTime
				fmt.Printf("Sleeping for %.2f (s)\n", timeDiff.Seconds())
				time.Sleep(timeDiff)

			}
		}
		println("Get")
		// идём получать данные тут один сайт не прописал 2 рой
		_, err := http.Get("http://www.example.com/index.html")
		if err != nil {
			panic(err)
		}
		// записываем когда обновляли данные 1 цикл пропускает все предидущие записаи
		if i%2 == 0 {
			lastRequestTime["packtpub.com"] = time.Now()
		} else {
			lastRequestTime["example.com"] = time.Now()
		}
	}
}
