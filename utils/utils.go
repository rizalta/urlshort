package utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func IsDeadLink(url string) bool {
	client := &http.Client{
		Timeout: 4 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return true
	}
	defer resp.Body.Close()
	return resp.StatusCode >= 400
}

func GenerateRandomString(size int) string {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""
	for i := 0; i < size; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}

func CurrentBaseURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s", scheme, r.Host)
}
