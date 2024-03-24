package main

import "fmt"

func scarperSite(url string, statusChan chan map[string]string) {
	statusChan <- map[string]string{url: "done"}
}

func m() {
	//мапа с сайтами для парсинга
	siteStatus := map[string]string{
		"https://example.com/page1.html": "ready",
		"https://example.com/page2.html": "ready",
		"https://example.com/page3.html": "ready",
	}
	//канал связи горутин
	updatesChan := make(chan map[string]string)
	//считаем кол-вщ завершёных горутин
	numberCompleted := 0
	//меняем статус в карте на воркинг и запускаем горутины
	//внутри которых происходит обновление на статус доне
	//после одного круга горутина блокируеться и ждёт пока очистится мапа канал
	//канал небуферизированный храни только одно значение
	for site := range siteStatus {
		siteStatus[site] = "WORKING"
		go scarperSite(site, updatesChan)
	}
	// берёт даные из канала в цикле
	//как только дойдёт до конца закроет канал
	//если я правильно понял он просто обновляется каждый раз в ебсконечном цикле
	//из за этого необходимо закрыть канал
	//пока канал открыт будет идти цикл(может привести к дедлоку)
	for update := range updatesChan {
		for url, status := range update {
			siteStatus[url] = status
			numberCompleted++
		}
		//закрываем канал что бы не вызывать деадлок
		if numberCompleted == len(siteStatus) {
			close(updatesChan)
		}
		fmt.Println(update)
	}
}
