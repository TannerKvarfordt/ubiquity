package datastructures_test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/TannerKvarfordt/ubiquity/datastructures"
)

// Benchmark the performance of the Data function.
func BenchmarkData(b *testing.B) {
	m := datastructures.NewConcurrentMap[int, string]()
	for i := 0; i < 1000000; i++ {
		m.Set(i, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Data()
	}
}

// Benchmark the performance of single-threaded inserts on
// datastructures.concurrentMap where the key does not exist.
func BenchmarkInsertAbsent(b *testing.B) {
	m := datastructures.NewConcurrentMap[int, string]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(i, "val")
	}
}

// Benchmark the performance of single-threaded inserts on
// sync.Map where the key does not yet exist.
func BenchmarkInsertAbsentSyncMap(b *testing.B) {
	m := sync.Map{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(i, "val")
	}
}

// Benchmark the performance of single-threaded inserts on
// datastructures.concurrentMap where the key already exists.
func BenchmarkInsertPresent(b *testing.B) {
	const (
		key = "key"
		val = "val"
	)
	m := datastructures.NewConcurrentMap[string, string]()
	m.Set(key, val)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(key, val)
	}
}

// Benchmark the performance of single-threaded inserts on
// sync.Map where the key already exists.
func BenchmarkInsertPresentSyncMap(b *testing.B) {
	const (
		key = "key"
		val = "val"
	)
	m := sync.Map{}
	m.Store(key, val)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(key, val)
	}
}

// Benchmark the performance of single-threaded inserts on
// datastructures.concurrentMap where the key already exists
// but maps to a different value than what is being inserted.
func BenchmarkInsertDifferent(b *testing.B) {
	const key = "key"
	m := datastructures.NewConcurrentMap[string, int]()
	m.Set(key, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(key, i)
	}
}

// Benchmark the performance of single-threaded inserts on
// sync.Map where the key already exists but maps to a
// different value than what is being inserted.
func BenchmarkInsertDifferentSyncMap(b *testing.B) {
	const key = "key"
	m := sync.Map{}
	m.Store(key, -1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(key, i)
	}
}

// Benchmark the performance of multi-threaded inserts on
// datastructures.concurrentMap where the key does not exist.
func BenchmarkMultiInsertAbsent(b *testing.B) {
	m := datastructures.NewConcurrentMap[int, string]()
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(key int) {
			m.Set(key, "val")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Benchmark the performance of multi-threaded inserts on
// sync.Map where the key does not yet exist.
func BenchmarkMultiInsertAbsentSyncMap(b *testing.B) {
	m := sync.Map{}
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(key int) {
			m.Store(key, "val")
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Benchmark the performance of multi-threaded inserts on
// datastructures.concurrentMap where the key already exists.
func BenchmarkMultiInsertPresent(b *testing.B) {
	const (
		key = "key"
		val = "val"
	)
	m := datastructures.NewConcurrentMap[string, string]()
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			m.Set(key, val)
			wg.Done()
		}()
	}
	wg.Wait()
}

// Benchmark the performance of multi-threaded inserts on
// sync.Map where the key already exists.
func BenchmarkMultiInsertPresentSyncMap(b *testing.B) {
	const (
		key = "key"
		val = "val"
	)
	m := sync.Map{}
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			m.Store(key, val)
			wg.Done()
		}()
	}
	wg.Wait()
}

// Benchmark the performance of multi-threaded inserts on
// datastructures.concurrentMap where the key already exists,
// but maps to a different value than what is being inserted.
func BenchmarkMultiInsertDifferent(b *testing.B) {
	const key = "key"
	m := datastructures.NewConcurrentMap[string, int]()
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(val int) {
			m.Set(key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// Benchmark the performance of multi-threaded inserts on
// sync.Map where the key already exists, but maps to a
// different value than what is being inserted.
func BenchmarkMultiInsertDifferentSyncMap(b *testing.B) {
	const key = "key"
	m := sync.Map{}
	wg := &sync.WaitGroup{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(val int) {
			m.Store(key, val)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
