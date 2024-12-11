package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse URL: %w", err)
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse URL: %w", err)
	}

	// Check if domains of base and current url are same, if not return current pages
	if baseURL.Host != currentURL.Host {
		return pages, nil
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't normalize URL: %w", err)
	}

	if _, ok := pages[normalizedCurrentURL]; ok {
		pages[normalizedCurrentURL] += 1
		return pages, nil
	} else {
		pages[normalizedCurrentURL] = 1
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't get html body: %w", err)
	}

	fmt.Printf("Crawling URL: %v\n", rawCurrentURL)

	urls, err := getURLsFromHTML(htmlBody, rawCurrentURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't get urls from current url: %w", err)
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

	return pages, nil
}
