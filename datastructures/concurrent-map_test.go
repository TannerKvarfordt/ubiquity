package datastructures_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/datastructures"
)

func TestConcurrentMap(t *testing.T) {
	m := datastructures.NewConcurrentMap[string, string]()
	if m.Len() != 0 {
		t.Error()
	}
	if _, ok := m.At(""); ok {
		t.Error()
	}
	if m.Len() != 0 {
		t.Error()
	}

	const (
		key1 = "SLATFATF"
		val1 = "So long, and thanks for all the fish."
	)
	m.Set(key1, val1)
	if m.Len() != 1 {
		t.Error()
	}
	if val, ok := m.At(key1); ok {
		if val != val1 {
			t.Error()
		}
	} else {
		t.Error()
	}

	const (
		key2 = val1
		val2 = key1
	)
	m.Set(key2, val2)
	if m.Len() != 2 {
		t.Error()
	}
	if val, ok := m.At(key2); ok {
		if val != val2 {
			t.Error()
		}
	} else {
		t.Error()
	}
}

func TestConcurrentMapSetIfPresent(t *testing.T) {
	m := datastructures.NewConcurrentMap[string, string]()

	const (
		key1 = "Douglas Adams"
		val1 = "So long, and thanks for all the fish."
		val2 = "Time is an illusion. Lunchtime doubly so."
	)

	if ok := m.SetIfPresent(key1, val2); ok {
		t.Error()
	}
	if _, ok := m.At(key1); ok {
		t.Error()
	}
	if m.Len() != 0 {
		t.Error()
	}

	m.Set(key1, val1)
	if ok := m.SetIfPresent(key1, val2); !ok {
		t.Error()
	}
	if val, ok := m.At(key1); ok {
		if val != val2 {
			t.Error()
		}
	} else {
		t.Error()
	}
	if m.Len() != 1 {
		t.Error()
	}
}

func TestConcurrentMapSetIfAbsent(t *testing.T) {
	m := datastructures.NewConcurrentMap[string, string]()

	const (
		key1 = "Douglas Adams"
		val1 = "So long, and thanks for all the fish."
		val2 = "Time is an illusion. Lunchtime doubly so."
	)

	if ok := m.SetIfAbsent(key1, val1); !ok {
		t.Error()
	}
	if val, ok := m.At(key1); ok {
		if val != val1 {
			t.Error()
		}
	} else {
		t.Error()
	}
	if m.Len() != 1 {
		t.Error()
	}

	if ok := m.SetIfAbsent(key1, val2); ok {
		t.Error()
	}
	if val, ok := m.At(key1); ok {
		if val != val1 {
			t.Error()
		}
	} else {
		t.Error()
	}
	if m.Len() != 1 {
		t.Error()
	}
}

func TestConcurrentMapData(t *testing.T) {
	const (
		key1 = "SLATFATF"
		val1 = "So long, and thanks for all the fish."

		key2 = "TIAILDS"
		val2 = "Time is an illusion. Lunchtime doubly so."
	)
	m := datastructures.NewConcurrentMap[string, string]()
	m.Set(key1, val1)
	m.Set(key2, val2)

	data := m.Data()
	if len(data) != 2 {
		t.Error()
	}
	m.Set(val1, key1)
	if len(data) != 2 {
		t.Error()
	}

	if data[key1] != val1 {
		t.Error()
	}
	if data[key2] != val2 {
		t.Error()
	}
}

func TestConcurrentMapRemove(t *testing.T) {
	const (
		key1 = "SLATFATF"
		val1 = "So long, and thanks for all the fish."

		key2 = "TIAILDS"
		val2 = "Time is an illusion. Lunchtime doubly so."
	)

	m := datastructures.NewConcurrentMap[string, string]()
	if removed := m.Remove("foobarbaz"); removed {
		t.Error()
	}
	if m.Len() != 0 {
		t.Error()
	}

	m.Set(key1, val1)
	m.Set(key2, val2)
	if m.Len() != 2 {
		t.Error()
	}

	if removed := m.Remove(key1 + "foobarbaz"); removed {
		t.Error()
	}

	if removed := m.Remove(key1); !removed {
		t.Error()
	}
	if m.Len() != 1 {
		t.Error()
	}
	if val, ok := m.At(key2); ok {
		if val != val2 {
			t.Error()
		}
	} else {
		t.Error()
	}

	if removed := m.Remove(key2); !removed {
		t.Error()
	}
	if m.Len() != 0 {
		t.Error()
	}
}

// TODO: benchmark against an unsharded implementation
