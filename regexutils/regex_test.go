package regexutils_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/regexutils"
)

func TestAltGroupEmpty(t *testing.T) {
	if regexutils.BuildAltGroup() != "" {
		t.Error()
	}
}

func TestAltGroupSingle(t *testing.T) {
	if regexutils.BuildAltGroup("a") != "(a)" {
		t.Error()
	}
}

func TestAltGroupDouble(t *testing.T) {
	if regexutils.BuildAltGroup("a, b") != "(a, b)" {
		t.Error()
	}
}
