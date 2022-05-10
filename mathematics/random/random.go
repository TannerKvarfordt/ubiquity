package random

import (
	"fmt"
	"math/rand"
)

// Return a non-negative random number in the inclusive range [min, max].
// If max <= min, returns the maximum uint value and an error.
// TODO: Support negative ranges?
func Range(min, max uint64) (uint64, error) {
	if max <= min {
		return ^uint64(0), fmt.Errorf("max (%d) cannot be less than or equal to min (%d)", max, min)
	}
	return uint64(rand.Intn(int((max-min)+1))) + min, nil
}

// Returns true or false, randomly.
func RandomBoolean() bool {
	return rand.Int31()&0x01 == 0
}
