package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func makeGetCall() (string, error) {

	client := &http.Client{}

	token := os.Getenv("GITHUB_TOKEN")

	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Invalid status %d", resp.StatusCode)
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
