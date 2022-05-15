package sliceutils_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/sliceutils"
)

func TestAllItemsEqual(t *testing.T) {
	if sliceutils.AllItemsEqual(1, 2, 3, 4) {
		t.Error()
	}
	if sliceutils.AllItemsEqual(2.3, 3.0000) {
		t.Error()
	}
	if sliceutils.AllItemsEqual("a", "b") {
		t.Error()
	}
	if sliceutils.AllItemsEqual(true, false, true, false, true, true, false) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual[int]() {
		t.Error()
	}
	if !sliceutils.AllItemsEqual[string]() {
		t.Error()
	}
	if !sliceutils.AllItemsEqual[float32]() {
		t.Error()
	}
	if !sliceutils.AllItemsEqual[float64]() {
		t.Error()
	}
	if !sliceutils.AllItemsEqual[bool]() {
		t.Error()
	}
	if !sliceutils.AllItemsEqual("a") {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(1) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(2.3) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(true) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual("a", "a") {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(1, 1) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(1.1, 1.1) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(false, false) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(true, true) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual("a", "a", "a") {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(1, 1, 1) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(1.1, 1.1, 1.1) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(false, false, false) {
		t.Error()
	}
	if !sliceutils.AllItemsEqual(true, true, true) {
		t.Error()
	}
}

func TestAllItemsUnique(t *testing.T) {
	if !sliceutils.AllItemsUnique("a") {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(1) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(2.3) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(true) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(1, 2) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(2.3, 3.0000) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique("a", "b") {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(true, false) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(1, 2, 3) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(2.3, 3.0001, 4.02) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique("a", "b", "c") {
		t.Error()
	}
	if !sliceutils.AllItemsUnique[int]() {
		t.Error()
	}
	if !sliceutils.AllItemsUnique[string]() {
		t.Error()
	}
	if !sliceutils.AllItemsUnique[float32]() {
		t.Error()
	}
	if !sliceutils.AllItemsUnique[float64]() {
		t.Error()
	}
	if !sliceutils.AllItemsUnique[bool]() {
		t.Error()
	}
	if sliceutils.AllItemsUnique("a", "a") {
		t.Error()
	}
	if sliceutils.AllItemsUnique(1, 1) {
		t.Error()
	}
	if sliceutils.AllItemsUnique(1.1, 1.1) {
		t.Error()
	}
	if sliceutils.AllItemsUnique(false, false) {
		t.Error()
	}
	if sliceutils.AllItemsUnique(true, true) {
		t.Error()
	}
	if sliceutils.AllItemsUnique("a", "a", "a") {
		t.Error()
	}
	if sliceutils.AllItemsUnique(1, 1, 1) {
		t.Error()
	}
	if sliceutils.AllItemsUnique(1.1, 1.1, 1.1) {
		t.Error()
	}
	if sliceutils.AllItemsUnique(false, false, false) {
		t.Error()
	}
	if sliceutils.AllItemsUnique(true, true, true) {
		t.Error()
	}
}

func TestEqual(t *testing.T) {
	if sliceutils.Equal([]int{}, []int{1}) {
		t.Error()
	}
	if sliceutils.Equal([]int{1}, []int{}) {
		t.Error()
	}
	if sliceutils.Equal([]int{1}, []int{1, 2}) {
		t.Error()
	}
	if sliceutils.Equal([]int{1, 2}, []int{1}) {
		t.Error()
	}
	if sliceutils.Equal([]int{1, 2}, []int{2, 1}) {
		t.Error()
	}
	if sliceutils.Equal([]int{1, 2, 3}, []int{3, 2, 1}) {
		t.Error()
	}

	if !sliceutils.Equal([]int{}, []int{}) {
		t.Error()
	}
	if !sliceutils.Equal([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}) {
		t.Error()
	}
	if !sliceutils.Equal([]string{"foo", "bar"}, []string{"foo", "bar"}) {
		t.Error()
	}
	if !sliceutils.Equal([]float32{2.0}, []float32{2.0}) {
		t.Error()
	}
	if !sliceutils.Equal([]bool{true, true, true, false}, []bool{true, true, true, false}) {
		t.Error()
	}

}

func TestRemoveDuplicates(t *testing.T) {
	// Test that order is unchanged
	if !sliceutils.Equal(sliceutils.RemoveDuplicates(1, 2, 3, 4, 1), []int{1, 2, 3, 4}) {
		t.Error()
	}
	if !sliceutils.Equal(sliceutils.RemoveDuplicates("foo", "bar", "baz"), []string{"foo", "bar", "baz"}) {
		t.Error()
	}
	if !sliceutils.Equal(sliceutils.RemoveDuplicates(4, 4, 4, 1, 1, 2, 3, 5, 5), []int{4, 1, 2, 3, 5}) {
		t.Error()
	}

	if !sliceutils.AllItemsUnique(sliceutils.RemoveDuplicates(1, 2, 44, 121, 1, 1, 5, 12, 121, 110, 100, 2, 1000)...) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(sliceutils.RemoveDuplicates("a", "b", "c", "q", "a", "q")...) {
		t.Error()
	}
	if !sliceutils.AllItemsUnique(sliceutils.RemoveDuplicates(2.1, 2.2, 2.1, 2.4, 5.11, 3.0, 32.23, 5.11, 2.1)...) {
		t.Error()
	}
}

func TestContains(t *testing.T) {
	tList := []string{"a", "b", "c"}
	if !sliceutils.Contains("a", tList...) {
		t.Error()
	}
	if !sliceutils.Contains("b", tList...) {
		t.Error()
	}
	if !sliceutils.Contains("c", tList...) {
		t.Error()
	}
	if sliceutils.Contains("d", tList...) {
		t.Error()
	}

	// No list provided, so does not contain keyVal.
	if sliceutils.Contains(1) {
		t.Error()
	}
	if sliceutils.Contains(1.2) {
		t.Error()
	}
	if sliceutils.Contains(-1) {
		t.Error()
	}
	if sliceutils.Contains(-1.3) {
		t.Error()
	}
	if sliceutils.Contains(true) {
		t.Error()
	}
	if sliceutils.Contains("foo") {
		t.Error()
	}
}

func TestCount(t *testing.T) {
	tList := []string{"a", "b", "c", "b", "c", "c"}
	if sliceutils.Count("a", tList...) != 1 {
		t.Error()
	}
	if sliceutils.Count("b", tList...) != 2 {
		t.Error()
	}
	if sliceutils.Count("c", tList...) != 3 {
		t.Error()
	}
	if sliceutils.Count("d", tList...) != 0 {
		t.Error()
	}

	// No list provided, so does not contain keyVal.
	if sliceutils.Count(1) != 0 {
		t.Error()
	}
	if sliceutils.Count(1.2) != 0 {
		t.Error()
	}
	if sliceutils.Count(-1) != 0 {
		t.Error()
	}
	if sliceutils.Count(-1.3) != 0 {
		t.Error()
	}
	if sliceutils.Count(true) != 0 {
		t.Error()
	}
	if sliceutils.Count("foo") != 0 {
		t.Error()
	}
}
