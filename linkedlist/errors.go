package linkedlist

import "fmt"

type IndexOutOfRangeError struct {
	idx int
}

func (e *IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("Index (%d) is out of range", e.idx)
}
