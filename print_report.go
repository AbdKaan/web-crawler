package main

import "fmt"

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	

	for x, y := range pages {
		fmt.Printf("Found %v internal links to %v\n", x, y)
	}
}
