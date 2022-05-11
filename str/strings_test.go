package str_test

import (
	"testing"
	"unicode/utf8"

	"github.com/TannerKvarfordt/ubiquity/str"
)

func TestFirstN(t *testing.T) {
	foostr := ""
	if str.FirstN(foostr, 0) != "" {
		t.Error()
	}
	if str.FirstN(foostr, 1) != "" {
		t.Error()
	}
	if str.FirstN(foostr, 2) != "" {
		t.Error()
	}

	foostr = "foo"
	if str.FirstN(foostr, 0) != "" {
		t.Error()
	}
	if str.FirstN(foostr, 1) != "f" {
		t.Error()
	}
	if str.FirstN(foostr, 2) != "fo" {
		t.Error()
	}
	if str.FirstN(foostr, uint64(utf8.RuneCountInString(foostr))) != foostr {
		t.Error()
	}
	if str.FirstN(foostr, uint64(utf8.RuneCountInString(foostr)+1)) != foostr {
		t.Error()
	}
	if str.FirstN(foostr, uint64(utf8.RuneCountInString(foostr)+2)) != foostr {
		t.Error()
	}

	foostr = "世界 hello"
	if str.FirstN(foostr, 0) != "" {
		t.Error()
	}
	if str.FirstN(foostr, 1) != "世" {
		t.Error()
	}
	if str.FirstN(foostr, 2) != "世界" {
		t.Error()
	}
	if str.FirstN(foostr, 3) != "世界 " {
		t.Error()
	}
	if str.FirstN(foostr, 4) != "世界 h" {
		t.Error()
	}
	if str.FirstN(foostr, uint64(utf8.RuneCountInString(foostr))) != foostr {
		t.Error()
	}
	if str.FirstN(foostr, uint64(utf8.RuneCountInString(foostr)+1)) != foostr {
		t.Error()
	}
	if str.FirstN(foostr, uint64(utf8.RuneCountInString(foostr)+2)) != foostr {
		t.Error()
	}
}
