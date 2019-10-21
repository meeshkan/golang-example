package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func makeGetCall() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}

func main() {
	makeGetCall()
}
