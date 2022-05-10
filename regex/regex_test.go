package regex_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/regex"
)

func TestAltGroupEmpty(t *testing.T) {
	if regex.BuildAltGroup() != "" {
		t.Error()
	}
}

func TestAltGroupSingle(t *testing.T) {
	if regex.BuildAltGroup("a") != "(a)" {
		t.Error()
	}
}

func TestAltGroupDouble(t *testing.T) {
	if regex.BuildAltGroup("a, b") != "(a, b)" {
		t.Error()
	}
}

func TestIsNotNumeric(t *testing.T) {
	t.Error("unimplemented")
}

func TestIsNumeric(t *testing.T) {
	t.Error("unimplemented")
}

func TestSentenceEnd(t *testing.T) {
	t.Error("unimplemented")
}
