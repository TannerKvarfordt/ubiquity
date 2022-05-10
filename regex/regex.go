package regex

import "regexp"

// Returns a regexp alternate group of the provided
// strings. For example, input of [a ,b] would result
// in a return value of "(a|b)". Calling this function with
// no arguments results in the empty string.
func BuildAltGroup(alts ...string) string {
	if len(alts) == 0 {
		return ""
	}

	altGroup := "("
	for i, alt := range alts {
		altGroup += alt
		if i+1 < len(alts) {
			altGroup += "|"
		}
	}
	altGroup += ")"
	return altGroup
}

var (
	IsNotNumericRegex = func() *regexp.Regexp { return nil }
	IsNumericRegex    = func() *regexp.Regexp { return nil }
)

func init() {
	r1 := regexp.MustCompile("[^0-9]+")
	IsNotNumericRegex = func() *regexp.Regexp { return r1 }

	r2 := regexp.MustCompile(`^\d+$`)
	IsNumericRegex = func() *regexp.Regexp { return r2 }
}

var SentenceEndPunctRegex = func() *regexp.Regexp { return nil }

func init() {
	r := regexp.MustCompile(`\s*[^\d\w>]+\s*$`)
	SentenceEndPunctRegex = func() *regexp.Regexp { return r }
}
