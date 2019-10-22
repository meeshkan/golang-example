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

// Repository is created from a GitHub repository object.
// See the documentation at https://developer.github.com/v3/repos/
type Repository struct {
	Name  string `json:"name"`
	Stars int    `json:"stargazers_count"`
}

func init() {
	if err := godotenv.Load(); err != nil {
		// Ignore
	}
}

func parseRepositories(body []byte) ([]Repository, error) {
	var parsed []Repository
	err := json.Unmarshal([]byte(body), &parsed)

	return parsed, err
}

func getGitHubRepositories() ([]Repository, error) {

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

	// TODO Pagination
	// link := resp.Header.Get("Link")
	// log.Println(link)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Invalid status %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseRepositories(body)
}

func computeStars(repositories *[]Repository) int {
	totalStars := 0
	for _, repository := range *repositories {
		totalStars += repository.Stars
	}
	return totalStars
}

func main() {
	repositories, err := getGitHubRepositories()
	if err != nil {
		log.Fatalln("Failed fetching repositories")
	}

	totalStars := computeStars(&repositories)

	log.Printf("You have %d repositories and %d stars", len(repositories), totalStars)

	// reposString, err := json.Marshal(repositories)
	// log.Println(string(reposString))
}
