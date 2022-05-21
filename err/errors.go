package err

import "fmt"

type IndexOutOfRange struct {
	Idx int
}

func (e *IndexOutOfRange) Error() string {
	return fmt.Sprintf("Index (%d) is out of range", e.Idx)
}
