package main

import (
	"net/http"
	"net/url"
	"log"
)


func Proxy() *http.Client {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:1080")
	}
	transport := &http.Transport{Proxy: proxy}

	return &http.Client{Transport: transport}
}

func main() {
	var debug bool
	token, err := GetBotToken("./token.json")

	if err != nil {
		log.Fatal(err)
	}

	robot := newRobot(token,
		"ZenJingBot", "", Proxy())
	robot.bot.Debug = debug
	robot.run()
}

