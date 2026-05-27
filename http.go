package main

import (
	"errors"
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
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/26.4 Safari/605.1.15")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("Status code " + resp.Status)
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
