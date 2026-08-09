package main

import (
	"log"
	"fmt"
	"flag"
)

var appVersion = "1.0.0"

func main() {
	printAbout := flag.Bool("about", false, "Print information about this app.")

	flag.Parse();

	if *printAbout {
		author := "paulie-of-punskas, a.k.a. paulie-aus-punskas"
		appAbout := "This is a CLI tool, that is used for monitoring costs of GitHub Actions runners" +
			"\nNo data is transmitted outside of URLs defined by user." +
			"\nNo AI data center was used to produce this code."
		log.Printf("[Author]\n%s \n\n[App version]\n%s \n\n[About]\n%s", author, appVersion, appAbout)
	} else {
		fmt.Printf("Available flags: %s\n", "\n-about")
	}
}
