package main

import (
	"io/ioutil"
	"encoding/json"
	"unicode"
)

func GetBotToken(path string) (string, error) {
	var data map[string]string

	file, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return "", err
	}

	return data["token"], nil
}


func chinese(words string) (zh bool) {
	for _, r := range words {
		if unicode.Is(unicode.Scripts["Han"], r) {
			zh = true
			break
		}
	}
	return
}