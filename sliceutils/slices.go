package sliceutils

// AllItemsEqual compares all provided elements and returns
// true if they are all equivalent.
func AllItemsEqual[T comparable](items ...T) bool {
	if len(items) == 0 {
		return true
	}

	first := items[0]
	for _, e := range items {
		if first != e {
			return false
		}
	}
	return true
}

// AllItemsUnique compares all provided elements
// and returns true if every element is unique.
func AllItemsUnique[T comparable](items ...T) bool {
	if len(items) == 0 {
		return true
	}
	set := make(map[T]bool, len(items))
	for _, e := range items {
		if _, ok := set[e]; ok {
			return false
		} else {
			set[e] = true
		}
	}
	return true
}

// RemoveDuplicates removes all duplicates from the provided elements,
// and returns a slice containing the unique set of elements.
// Order is preserved.
func RemoveDuplicates[T comparable](items ...T) []T {
	keyset := make(map[T]bool, len(items))
	rv := make([]T, 0, len(items))
	for _, item := range items {
		if _, ok := keyset[item]; !ok {
			keyset[item] = true
			rv = append(rv, item)
		}
	}
	return rv
}

// Equal compares two slices. If they are the same length,
// and all of their elements are identical and appear in the
// same order, they are considered equal.
func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
