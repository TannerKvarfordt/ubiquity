package regexutils

// Returns a regex alternate group of the provided
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
