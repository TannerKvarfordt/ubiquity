package mathutils

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](first T, args ...T) T {
	min := first
	for _, t := range args {
		if t < min {
			min = t
		}
	}
	return min
}
