package sorting_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/sorting"
)

func TestSlicesEqual(t *testing.T) {
	a := []int{1, 2, 3}
	if sorting.SlicesEqual(a, nil) {
		t.Fatal()
	}
	if sorting.SlicesEqual(nil, a) {
		t.Fatal()
	}

	b := []int{1, 2}
	if sorting.SlicesEqual(a, b) {
		t.Fatal()
	}
	if sorting.SlicesEqual(b, a) {
		t.Fatal()
	}

	b = []int{3, 1, 2}
	if sorting.SlicesEqual(a, b) {
		t.Fatal()
	}
	if sorting.SlicesEqual(b, a) {
		t.Fatal()
	}

	copy(b, a)
	if !sorting.SlicesEqual(a, b) {
		t.Fatal()
	}
	if !sorting.SlicesEqual(b, a) {
		t.Fatal()
	}

	if !sorting.SlicesEqual[int](nil, nil) {
		t.Fatal()
	}

	if !sorting.SlicesEqual([]int{}, []int{}) {
		t.Fatal()
	}
}
