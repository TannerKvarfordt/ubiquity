package mathutils_test

import (
	"github.com/TannerKvarfordt/ubiquity/mathutils"
	"testing"
)

// -------- Ints -------- //
func TestMinIntBack(t *testing.T) {
	// int8
	{
		nums := []int8{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int16
	{
		nums := []int16{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int32
	{
		nums := []int32{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int64
	{
		nums := []int64{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int
	{
		nums := []int{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
}
func TestMinIntMid(t *testing.T) {
	// int8
	{
		nums := []int8{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int16
	{
		nums := []int16{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int32
	{
		nums := []int32{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int64
	{
		nums := []int64{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int
	{
		nums := []int{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
}
func TestMinIntFront(t *testing.T) {
	// int8
	{
		nums := []int8{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int16
	{
		nums := []int16{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int32
	{
		nums := []int32{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int64
	{
		nums := []int64{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// int
	{
		nums := []int{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
}

func TestMinIntNegative(t *testing.T) {
	// One negative
	{
		nums := []int{1, -2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != -2 {
			t.Error()
		}
	}

	// two negatives
	{
		nums := []int{-1, -2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != -2 {
			t.Error()
		}
	}

	// all negative
	{
		nums := []int{-1, -2, -3, -4}
		if mathutils.Min(nums[0], nums[1:]...) != -4 {
			t.Error()
		}
	}

	// 0
	{
		nums := []int{0, -2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != -2 {
			t.Error()
		}
	}
}

func TestMinIntDuplicates(t *testing.T) {
	// one duplicate
	{
		nums := []int{1, 1, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}

	// two duplicates
	{
		nums := []int{1, 1, 4, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}

	// one entry
	{
		nums := []int{4}
		if mathutils.Min(nums[0]) != 4 {
			t.Error()
		}
	}

	// all duplicates
	{
		nums := []int{4, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 4 {
			t.Error()
		}
	}
}

// -------- Uints -------- //
func TestMinUintBack(t *testing.T) {
	// uint8
	{
		nums := []uint8{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint16
	{
		nums := []uint16{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint32
	{
		nums := []uint32{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint64
	{
		nums := []uint64{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint
	{
		nums := []uint{4, 3, 2, 1}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
}
func TestMinUintMid(t *testing.T) {
	// uint8
	{
		nums := []uint8{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint16
	{
		nums := []uint16{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint32
	{
		nums := []uint32{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint64
	{
		nums := []uint64{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint
	{
		nums := []uint{2, 1, 4, 3}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
}
func TestMinUintFront(t *testing.T) {
	// uint8
	{
		nums := []uint8{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint16
	{
		nums := []uint16{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint32
	{
		nums := []uint32{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint64
	{
		nums := []uint64{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
	// uint
	{
		nums := []uint{1, 2, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}
}

func TestMinUintDuplicates(t *testing.T) {
	// one duplicate
	{
		nums := []uint{1, 1, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}

	// two duplicates
	{
		nums := []uint{1, 1, 4, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}

	// one entry
	{
		nums := []uint{4}
		if mathutils.Min(nums[0]) != 4 {
			t.Error()
		}
	}

	// all duplicates
	{
		nums := []uint{4, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 4 {
			t.Error()
		}
	}
}

// -------- Floats -------- //
func TestMinFloatBack(t *testing.T) {
	// float32
	{
		nums := []float32{4.0, 3.0, 2.0, 1.0}
		if mathutils.Min(nums[0], nums[1:]...) != 1.0 {
			t.Error()
		}
	}
	// float64
	{
		nums := []float64{4.0, 3.0, 2.0, 1.0}
		if mathutils.Min(nums[0], nums[1:]...) != 1.0 {
			t.Error()
		}
	}
}
func TestMinFloatMid(t *testing.T) {
	// float32
	{
		nums := []float32{3.0, 1.0, 4.0, 2.0}
		if mathutils.Min(nums[0], nums[1:]...) != 1.0 {
			t.Error()
		}
	}
	// float64
	{
		nums := []float64{3.0, 1.0, 4.0, 2.0}
		if mathutils.Min(nums[0], nums[1:]...) != 1.0 {
			t.Error()
		}
	}
}
func TestMinFloatFront(t *testing.T) {
	// float32
	{
		nums := []float32{1.0, 3.0, 2.0, 4.0}
		if mathutils.Min(nums[0], nums[1:]...) != 1.0 {
			t.Error()
		}
	}
	// float64
	{
		nums := []float64{1.0, 3.0, 2.0, 4.0}
		if mathutils.Min(nums[0], nums[1:]...) != 1.0 {
			t.Error()
		}
	}
}

func TestMinFloatNegative(t *testing.T) {
	// One negative
	{
		nums := []float32{1.0, -2.0, 3.0, 4.0}
		if mathutils.Min(nums[0], nums[1:]...) != -2.0 {
			t.Error()
		}
	}

	// two negatives
	{
		nums := []float32{-1.0, -2.0, 3.0, 4.0}
		if mathutils.Min(nums[0], nums[1:]...) != -2.0 {
			t.Error()
		}
	}

	// all negative
	{
		nums := []float32{-1.0, -2.0, -3.0, -4.0}
		if mathutils.Min(nums[0], nums[1:]...) != -4.0 {
			t.Error()
		}
	}

	// 0
	{
		nums := []float32{0.0, -2.0, 3.0, 4.0}
		if mathutils.Min(nums[0], nums[1:]...) != -2.0 {
			t.Error()
		}
	}
}

func TestMinFloatDuplicates(t *testing.T) {
	// one duplicate
	{
		nums := []float64{1, 1, 3, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}

	// two duplicates
	{
		nums := []float64{1, 1, 4, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 1 {
			t.Error()
		}
	}

	// one entry
	{
		nums := []float64{4}
		if mathutils.Min(nums[0]) != 4 {
			t.Error()
		}
	}

	// all duplicates
	{
		nums := []float64{4, 4}
		if mathutils.Min(nums[0], nums[1:]...) != 4 {
			t.Error()
		}
	}
}

// -------- Strings -------- //
func TestMinStringBack(t *testing.T) {
	stringutilss := []string{"d", "c", "b", "a"}
	if mathutils.Min(stringutilss[0], stringutilss[1:]...) != "a" {
		t.Error()
	}
}
func TestMinStringMid(t *testing.T) {
	stringutilss := []string{"c", "a", "d", "b"}
	if mathutils.Min(stringutilss[0], stringutilss[1:]...) != "a" {
		t.Error()
	}
}
func TestMinStringFront(t *testing.T) {
	stringutilss := []string{"a", "b", "c", "d"}
	if mathutils.Min(stringutilss[0], stringutilss[1:]...) != "a" {
		t.Error()
	}
}

func TestMinStringDuplicates(t *testing.T) {
	// one duplicate
	{
		stringutilss := []string{"a", "a", "c", "d"}
		if mathutils.Min(stringutilss[0], stringutilss[1:]...) != "a" {
			t.Error()
		}
	}

	// two duplicates
	{
		stringutilss := []string{"a", "a", "d", "d"}
		if mathutils.Min(stringutilss[0], stringutilss[1:]...) != "a" {
			t.Error()
		}
	}

	// one entry
	{
		stringutilss := []string{"d"}
		if mathutils.Min(stringutilss[0]) != "d" {
			t.Error()
		}
	}

	// all duplicates
	{
		stringutilss := []string{"d", "d"}
		if mathutils.Min(stringutilss[0]) != "d" {
			t.Error()
		}
	}
}
