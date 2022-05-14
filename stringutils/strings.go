package stringutils

// Returns the first n characters of a string.
func FirstN(s string, n uint64) string {
	i := uint64(0)
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}
