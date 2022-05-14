package httputils

import (
	"net/http"
	"strings"
)

func IsReachableURL(url string) bool {
	_, err := http.Get(url)
	return err == nil
}

func IsHTTPS(url string) bool {
	resp, err := http.Get(url)
	if err != nil {
		return false
	}
	// final URL resolved
	return strings.HasPrefix(resp.Request.URL.String(), "https://")
}
