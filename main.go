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
	fmt.Printf("starting crawl of: %v\n", os.Args[1])
	body, err := getHTML(os.Args[1]) 
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(body)
	os.Exit(0)
}
