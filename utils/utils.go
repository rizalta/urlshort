package utils

import (
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
)

func EnforceHTTP(URL string) string {
	if len(URL) < 4 {
		URL = "https://" + URL
	} else if URL[:4] != "http" {
		URL = "https://" + URL
	}
	return URL
}

func IsValidURL(URL string) bool {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}
	_, err = net.LookupHost(u.Host)
	if err != nil {
		return false
	}
	return true
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
