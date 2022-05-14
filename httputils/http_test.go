package httputils_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/httputils"
)

func TestIsReachable(t *testing.T) {
	url := ""
	if httputils.IsReachableURL(url) {
		t.Error()
	}

	url = "not a valid url"
	if httputils.IsReachableURL(url) {
		t.Error()
	}

	url = "https://www.github.com/TannerKvarfordt/ubiquity/httputils"
	if !httputils.IsReachableURL(url) {
		t.Error()
	}

	url = "https://www.google.com/"
	if !httputils.IsReachableURL(url) {
		t.Error()
	}

	url = "http://www.google.com/"
	if !httputils.IsReachableURL(url) {
		t.Error()
	}
}

func TestIsHttps(t *testing.T) {
	url := ""
	if httputils.IsHTTPS(url) {
		t.Error()
	}

	url = "not a valid url"
	if httputils.IsHTTPS(url) {
		t.Error()
	}

	url = "https://www.github.com/TannerKvarfordt/ubiquity/httputils"
	if !httputils.IsHTTPS(url) {
		t.Error()
	}

	url = "https://www.google.com/"
	if !httputils.IsHTTPS(url) {
		t.Error()
	}

	url = "http://www.google.com/"
	if httputils.IsHTTPS(url) {
		t.Error()
	}
}
