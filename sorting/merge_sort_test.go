// This file only includes tests for Merge -- MergeSort tests can be found in sort.go
package sorting_test

import (
	"github.com/TannerKvarfordt/ubiquity/sorting"
	"testing"
)

func TestMergeSort(t *testing.T) {
	if !sorting.SlicesEqual(sorting.MergeSort[int](), nil) {
		t.Error()
	}

	if !sorting.SlicesEqual(sorting.MergeSort(1), []int{1}) {
		t.Error()
	}

	if !sorting.SlicesEqual(sorting.MergeSort(1, 2, 3, 4, 5), []int{1, 2, 3, 4, 5}) {
		t.Error()
	}

	if !sorting.SlicesEqual(sorting.MergeSort(5, 4, 3, 2, 1), []int{1, 2, 3, 4, 5}) {
		t.Error()
	}

	if !sorting.SlicesEqual(sorting.MergeSort(5, 3, 1, 4, 2), []int{1, 2, 3, 4, 5}) {
		t.Error()
	}
}

func TestSortedMerge(t *testing.T) {
	if !sorting.SlicesEqual(sorting.Merge[int](nil, nil), []int{}) {
		t.Error()
	}

	a := []int{1}
	if !sorting.SlicesEqual(sorting.Merge(a, nil), a) {
		t.Error()
	}
	if !sorting.SlicesEqual(sorting.Merge(nil, a), a) {
		t.Error()
	}

	b := []int{1}
	if !sorting.SlicesEqual(sorting.Merge(a, b), append(a, b...)) {
		t.Error()
	}
	if !sorting.SlicesEqual(sorting.Merge(b, a), append(a, b...)) {
		t.Error()
	}

	b = []int{2}
	if !sorting.SlicesEqual(sorting.Merge(a, b), append(a, b...)) {
		t.Error()
	}
	if !sorting.SlicesEqual(sorting.Merge(b, a), append(a, b...)) {
		t.Error()
	}

	b = []int{2, 3, 4}
	if !sorting.SlicesEqual(sorting.Merge(a, b), append(a, b...)) {
		t.Error()
	}
	if !sorting.SlicesEqual(sorting.Merge(b, a), append(a, b...)) {
		t.Error()
	}

	a = []int{1, 17, 19}
	b = []int{4, 8, 12}
	if !sorting.SlicesEqual(sorting.Merge(a, b), []int{1, 4, 8, 12, 17, 19}) {
		t.Error()
	}
	if !sorting.SlicesEqual(sorting.Merge(b, a), []int{1, 4, 8, 12, 17, 19}) {
		t.Error()
	}
}

func TestUnsortedMerged(t *testing.T) {
	a := []int{4, 1, 3}
	b := []int{6, 2}
	if !sorting.SlicesEqual(sorting.Merge(a, b), []int{4, 1, 3, 6, 2}) {
		t.Error()
	}
}