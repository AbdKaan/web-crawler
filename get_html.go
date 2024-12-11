package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't get result: %w", err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 399 {
		return "", fmt.Errorf("response failed with status code: %d and body: %s", res.StatusCode, body)
	}
	if strings.Contains(res.Header.Get("content-type"), "test/html") {
		return "", fmt.Errorf("header content-type is not text/html, it is: %v", res.Header.Get("content-type"))
	}
	if err != nil {
		return "", fmt.Errorf("couldn't read body: %w", err)
	}

	return string(body), nil
}
