package main

import (
	"io"
	"net/http"
	"strings"
	"time"
)

var defaultClient = &http.Client{Timeout: 10 * time.Second}

func get(url string, client *http.Client) (string, error) {
	if client == nil {
		client = defaultClient
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "crawl")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return "", errorf("status code %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func cleanUrl(u string) string {
	if before, _, ok := strings.Cut(u, "?"); ok {
		return before
	}
	return u
}
