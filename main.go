package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// Expected command prompt, continue
	url := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", url)

	pages := make(map[string]int)
	pages, err := crawlPage(url, url, pages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for url, amount := range pages {
		fmt.Printf("URL: %s - Amount: %v\n", url, amount)
	}

	os.Exit(0)
}
