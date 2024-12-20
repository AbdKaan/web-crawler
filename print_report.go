package main

import (
    "fmt"
    "sort"
)

type Page struct {
    URL   string
    Count int
}

func printReport(pages map[string]int, baseURL string) {
    fmt.Println("=============================")
    fmt.Printf("REPORT for %s\n", baseURL)
    fmt.Println("=============================")

    sortedPages := sortPages(pages)

    for _, page := range sortedPages {
        fmt.Printf("Found %v internal links to %v\n", page.Count, page.URL)
    }
}

func sortPages(pages map[string]int) []Page {
    pagesSlice := []Page{}
    for URL, count := range pages {
        pagesSlice = append(pagesSlice, Page{URL: URL, Count: count})
    }

    sort.Slice(pagesSlice, func(i, j int) bool {
        if pagesSlice[i].Count == pagesSlice[j].Count {
            return pagesSlice[i].URL < pagesSlice[j].URL
        }
        return pagesSlice[i].Count > pagesSlice[j].Count
    })

    return pagesSlice
}