package main

import (
	// "flag"
	"fmt"
	"os"
)

var appVersion = "1.0.0"
var gitHubPat = ""

func main() {
	// set flags
	// aboutCmd := flag.NewFlagSet("about", flag.ExitOnError)
	// connectionCmd := flag.NewFlagSet("check_connection", flag.ExitOnError)
	// connectionRepoOwnerAndRepoName := connectionCmd.String("repoOwnerName", "", "Owner of the repository and repo name")

	if len(os.Args) == 1 {
		fmt.Printf("Available flags:\n%s%s", "-about\n", "-check_connection")
		os.Exit(0)
	}

	switch os.Args[1] {
	case "check_connection":
		// connectionCmd.Parse(os.Args[2:])
		userInput := os.Args[2]
		if len(userInput) > 0 {
			fmt.Printf("Checking if connection is possible for %s... ", userInput)
			if IsConnectionToGitHubPossible(userInput, false) {
				fmt.Println("Success!")
			}
		} else {
			fmt.Println("Please provide repository owner and repo name, e.g. 'paulie-of-punskas/gha-costs'")
		}
	case "about":
		author := "paulie-of-punskas, a.k.a. paulie-aus-punskas"
		appAbout := "This is a CLI tool, that is used for monitoring costs of GitHub Actions runners" +
			"\nNo data is transmitted outside of URLs defined by user." +
			"\nNo AI data center was used to produce this code."
		fmt.Printf("[Author]\n%s \n\n[App version]\n%s \n\n[About]\n%s\n", author, appVersion, appAbout)
	}
}
