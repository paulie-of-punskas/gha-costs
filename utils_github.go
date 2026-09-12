package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Function looks for GitHub PAT in following order:
// 1. .env, 2. environmental GITHUB_PAT variable
func GetPAT() string {
	envFileContent, err := os.ReadFile(".env")
	if err == nil {
		return string(envFileContent)
	}

	if patContent := os.Getenv("GITHUB_PAT"); patContent != "" {
		return patContent
	}

	fmt.Println("Neither .env, nor GITHUB_PAT were found.")
	return ""
}

// Using getPAT() results, check if GitHub API returns 200
// github PAT is optional - depends on how user had setup repository
func IsConnectionToGitHubPossible(repoOwnerAndName string, useGithubPAT bool) bool {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.github.com/repos/%s/actions/runs", repoOwnerAndName)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	if useGithubPAT {
		req.Header.Set("Authorization", "Bearer "+GetPAT())
	}
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Could not send GET request to %s: %s", repoOwnerAndName, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	}
	return false
}
