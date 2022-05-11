package networking_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/networking"
)

func TestIsReachable(t *testing.T) {
	url := ""
	if networking.IsReachableURL(url) {
		t.Error()
	}

	url = "not a valid url"
	if networking.IsReachableURL(url) {
		t.Error()
	}

	url = "https://www.github.com/TannerKvarfordt/ubiquity/networking"
	if !networking.IsReachableURL(url) {
		t.Error()
	}

	url = "https://www.google.com/"
	if !networking.IsReachableURL(url) {
		t.Error()
	}

	url = "http://www.google.com/"
	if !networking.IsReachableURL(url) {
		t.Error()
	}
}

func TestIsHttps(t *testing.T) {
	url := ""
	if networking.IsHTTPS(url) {
		t.Error()
	}

	url = "not a valid url"
	if networking.IsHTTPS(url) {
		t.Error()
	}

	url = "https://www.github.com/TannerKvarfordt/ubiquity/networking"
	if !networking.IsHTTPS(url) {
		t.Error()
	}

	url = "https://www.google.com/"
	if !networking.IsHTTPS(url) {
		t.Error()
	}

	url = "http://www.google.com/"
	if networking.IsHTTPS(url) {
		t.Error()
	}
}
