package mathematics

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](first T, args ...T) T {
	max := first
	for _, t := range args {
		if t > max {
			max = t
		}
	}
	return max
}
