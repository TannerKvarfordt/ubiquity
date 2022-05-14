package stringutils_test

import (
	"testing"
	"unicode/utf8"

	"github.com/TannerKvarfordt/ubiquity/stringutils"
)

func TestFirstN(t *testing.T) {
	foostringutils := ""
	if stringutils.FirstN(foostringutils, 0) != "" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 1) != "" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 2) != "" {
		t.Error()
	}

	foostringutils = "foo"
	if stringutils.FirstN(foostringutils, 0) != "" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 1) != "f" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 2) != "fo" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, uint64(utf8.RuneCountInString(foostringutils))) != foostringutils {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, uint64(utf8.RuneCountInString(foostringutils)+1)) != foostringutils {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, uint64(utf8.RuneCountInString(foostringutils)+2)) != foostringutils {
		t.Error()
	}

	foostringutils = "世界 hello"
	if stringutils.FirstN(foostringutils, 0) != "" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 1) != "世" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 2) != "世界" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 3) != "世界 " {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, 4) != "世界 h" {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, uint64(utf8.RuneCountInString(foostringutils))) != foostringutils {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, uint64(utf8.RuneCountInString(foostringutils)+1)) != foostringutils {
		t.Error()
	}
	if stringutils.FirstN(foostringutils, uint64(utf8.RuneCountInString(foostringutils)+2)) != foostringutils {
		t.Error()
	}
}
