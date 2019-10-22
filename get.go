package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeGetCall() (string, error) {

	client := &http.Client{}

	token := "fake"

	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.Status != "200" {
		return "", fmt.Errorf("Invalid status %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(body))
	return string(body), nil
}

func main() {
	makeGetCall()
}
