package sorting

import "golang.org/x/exp/constraints"

// SlicesEqual compares two slices, and returns true
// if they are the same length and compare equally
// at each index.
func SlicesEqual[T constraints.Ordered](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
