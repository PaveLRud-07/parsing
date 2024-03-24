package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"path"
)

var proxies []string = []string{
	"hhtp://31.186.239.244:8080",
	"hhtp://134.122.58.174:80",
	"hhtp://154.79.254.236:32650",
}

func GetProxy(_ *http.Request) (*url.URL, error) {
	randomIndex := rand.Int31n(int32(len(proxies)) - int32(1))
	randomProxy := proxies[randomIndex]
	return url.Parse(randomProxy)
}

func whitelist() {
	http.DefaultTransport.(*http.Transport).Proxy = GetProxy
	p, e := url.Parse("https://hub.packtpub.com/8-programming-languages-to-learn-in-2019/")
	if e != nil {
		panic(e)
	}
	site := p.Host + p.Path
	dM, er := path.Match("https://hub.packtpub.com/*", site)
	if er != nil {
		panic(er)
	}
	if dM {
		//continuescrap
	}
}
