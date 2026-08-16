package main

import (
	"fmt"
	"os"
)

// Function looks for GitHub PAT in following order:
// 1. .env, 2. env variable GITHUB_PAT
func GetPAT() string {

	if patContent := os.Getenv("GITHUB_PAT"); patContent != "" {
		return patContent
	}

	envFileContent, err := os.ReadFile(".env")
	if err == nil {
		return string(envFileContent)
	}

	fmt.Println("Neither .env, nor GITHUB_PAT were found.")
	return ""
}

// Using getPAT() results, check if GitHub API returns 200
func checkGitHubConnection() {
}
