package sorting

import "golang.org/x/exp/constraints"

type SortImpl uint

const (
	MergeImpl SortImpl = iota

	MergeImplEnd // Must always be last
)

// Sort returns provided args as a sorted list. If no args are provided,
// nil is returned. The sort is performed as defined by the provided SortImpl.
//
// TODO: provide a SortCustomComparator function which sorts based on a provided comparison function.
func Sort[T constraints.Ordered](impl SortImpl, args ...T) []T {
	switch impl {
	case MergeImpl:
		return MergeSort(args...)
	}
	return nil
}
