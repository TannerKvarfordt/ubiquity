package sorting

import "golang.org/x/exp/constraints"

// MergeSort returns provided args as a sorted list. If no args are provided,
// nil is returned. The sort is performed using the Merge Sort algorithm.
func MergeSort[T constraints.Ordered](args ...T) []T {
	switch len(args) {
	case 0:
		return nil
	case 1:
		return args
	}
	return Merge(MergeSort(args[0:len(args)/2]...), MergeSort(args[len(args)/2:]...))
}

// Merge returns a merged list comprised of two input lists.
// If the provided lists are sorted, the returned list will be as well.
func Merge[T constraints.Ordered](a, b []T) []T {
	merged := make([]T, len(a)+len(b))
	i, j, k := 0, 0, 0
	for i < len(a) && j < len(b) && k < len(merged) {
		if a[i] < b[j] {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
		k++
	}

	for i < len(a) {
		merged[k] = a[i]
		k++
		i++
	}
	for j < len(b) {
		merged[k] = b[j]
		k++
		j++
	}

	return merged
}
