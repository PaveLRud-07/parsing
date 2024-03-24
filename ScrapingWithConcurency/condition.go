package main

import (
	"sync"
	"time"
)

var sites []string = []string{
	"https://example.com/page1.html",
	"https://example.com/page2.html",
	"https://example.com/page3.html",
}
var activeThreads = 0
var doneCount = 0

const maxActiveThreads = 1

func scarperSite(site string, condition *sync.Cond) {
	condition.L.Lock()
	if activeThreads >= maxActiveThreads {
		//ждём сигнала автоматически разблокирует l.lock
		// автоматические его заблокируте после сигнала
		condition.Wait()
	}
	activeThreads++
	condition.L.Unlock()
	println("scraping" + site)
	//scarper code
	condition.L.Lock()
	activeThreads--
	doneCount++
	condition.L.Unlock()
	//сигнал конда
	condition.Signal()
}
func main() {
	var l = sync.Mutex{}
	var c = sync.NewCond(&l)
	for _, site := range sites {
		println("starting scarper for " + site)
		go scarperSite(site, c)
	}
	for doneCount < len(sites) {
		time.Sleep(1 * time.Second)
	}
}
