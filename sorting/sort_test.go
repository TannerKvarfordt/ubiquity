package sorting_test

import (
	"github.com/TannerKvarfordt/ubiquity/sorting"
	"testing"
)

func TestSort(t *testing.T) {
	for impl := sorting.SortImpl(0); impl < sorting.MergeImplEnd; impl++ {
		if !sorting.SlicesEqual(sorting.Sort(impl, "d", "a", "b", "c"), []string{"a", "b", "c", "d"}) {
			t.Error()
		}
	}
}
