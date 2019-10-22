package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		// Ignore
	}
}

func parseResponse(body []byte) ([]interface{}, error) {
	var parsed []interface{}
	err := json.Unmarshal([]byte(body), &parsed)

	return parsed, err
}

func getGitHubRepositories() ([]interface{}, error) {

	client := &http.Client{}

	token, exists := os.LookupEnv("GITHUB_TOKEN")

	if !exists {
		log.Fatalln("Environment variable GITHUB_TOKEN not set")
	}

	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("Authorization", fmt.Sprintf("token %s", token))

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseResponse(body)
}

func main() {
	repositories, err := getGitHubRepositories()
	if err != nil {
		log.Fatalln("Failed fetching repositories")
	}

	stars := 0
	for _, repository := range repositories {
		repo := repository.(map[string]interface{})
		stars += int(repo["stargazers_count"].(float64))
	}

	log.Printf("You have %d repositories and %d stars", len(repositories), stars)

	// reposString, err := json.Marshal(repositories)

	// log.Println(string(reposString))
}
