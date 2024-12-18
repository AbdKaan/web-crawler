package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// Expected command prompt, continue
	baseURL := os.Args[1]
	fmt.Printf("starting crawl of: %v\n", baseURL)

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("couldn't parse URL: %v", err)
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("max concurrency must be a number, it was: %v", err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("max pages must be a number, it was: %v", err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	// Start crawling
	cfg.wg.Add(1)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	// Print results
	/*
		for url, amount := range cfg.pages {
			fmt.Printf("URL: %s - Amount: %v\n", url, amount)
		}
	*/

	printReport(cfg.pages, baseURL)

	os.Exit(0)
}
