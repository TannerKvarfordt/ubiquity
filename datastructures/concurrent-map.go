package datastructures

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

const shardCount = 128

type shard[KEY comparable, VAL any] struct {
	Contents map[KEY]VAL
	mutex    sync.RWMutex
}

func newShard[KEY comparable, VAL any]() *shard[KEY, VAL] {
	return &shard[KEY, VAL]{
		Contents: map[KEY]VAL{},
	}
}

func (s *shard[KEY, VAL]) at(k KEY) (VAL, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	if v, ok := s.Contents[k]; ok {
		return v, ok
	}
	var v VAL
	return v, false
}

func (s *shard[KEY, VAL]) set(k KEY, v VAL) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.Contents[k] = v
}

func (s *shard[KEY, VAL]) setIfAbsent(k KEY, v VAL) {
	if _, ok := s.at(k); !ok {
		s.set(k, v)
		return
	}
}

func (s *shard[KEY, VAL]) setIfPresent(k KEY, v VAL) {
	if _, ok := s.at(k); !ok {
		return
	}
	s.set(k, v)
}

func (s *shard[_, VAL]) forEach(f func(v VAL) error, exitOnError bool) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var errs error = nil
	for i := range s.Contents {
		if err := f(s.Contents[i]); err != nil {
			if exitOnError {
				return err
			}
			if errs == nil {
				errs = err
			} else {
				errs = fmt.Errorf("%w : %s", errs, err)
			}
		}
	}
	return errs
}

// ConcurrentMap is a sharded, thread-safe mapping of keys to values.
type ConcurrentMap[KEY comparable, VAL any] struct {
	shards [shardCount]*shard[KEY, VAL]
}

// NewConcurrentMap creates a new ConcurrentMap.
func NewConcurrentMap[KEY comparable, VAL any]() *ConcurrentMap[KEY, VAL] {
	c := ConcurrentMap[KEY, VAL]{
		shards: [shardCount]*shard[KEY, VAL]{},
	}
	for i := range c.shards {
		c.shards[i] = newShard[KEY, VAL]()
	}
	return &c
}

func (c ConcurrentMap[KEY, VAL]) getShard(key KEY) *shard[KEY, VAL] {
	return c.shards[fnv1a(key)%uint64(shardCount)]
}

// At returns the value for the given key.
func (c *ConcurrentMap[KEY, VAL]) At(key KEY) (VAL, bool) {
	shard := c.getShard(key)
	return shard.at(key)
}

// Set inserts the given val at the specified key, or overrides the
// current value for that key if there is one.
func (c *ConcurrentMap[KEY, VAL]) Set(key KEY, val VAL) {
	shard := c.getShard(key)
	shard.set(key, val)
}

func (c *ConcurrentMap[KEY, VAL]) SetIfAbsent(key KEY, val VAL) {
	shard := c.getShard(key)
	shard.setIfAbsent(key, val)
}

func (c *ConcurrentMap[KEY, VAL]) SetIfPresent(key KEY, val VAL) {
	shard := c.getShard(key)
	shard.setIfPresent(key, val)
}

// Len returns the number of keys in the map.
func (c *ConcurrentMap[KEY, VAL]) Len() int {
	size := 0
	for _, s := range c.shards {
		if s == nil {
			continue
		}
		s.mutex.RLock()
		size += len(s.Contents)
		s.mutex.RUnlock()
	}
	return size
}

// ForEach runs the provided function for each element in the ConcurrentMap.
// If exitOnError is true, then this function exits as soon as an error is encountered.
// Otherwise, execution continues until the provided function is run on every element
// in the map, wrapping any errors encountered in the process.
func (c *ConcurrentMap[_, VAL]) ForEach(f func(val VAL) error, exitOnError bool) error {
	var errs error = nil
	for _, s := range c.shards {
		if err := s.forEach(f, exitOnError); err != nil {
			if exitOnError {
				return err
			}
			if errs == nil {
				errs = err
			} else {
				errs = fmt.Errorf("%w : %s", errs, err)
			}
		}
	}
	return errs
}

// https://en.wikipedia.org/wiki/Fowler-Noll-Vo_hash_function
func fnv1a[KEY comparable](key KEY) uint64 {
	const fnvOffsetBasis = uint64(0xcbf29ce484222325)
	const fnvPrime = uint64(0x100000001b3)

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(key); err != nil {
		panic(err)
	}

	hash := fnvOffsetBasis
	for _, b := range buf.Bytes() {
		hash ^= uint64(b)
		hash *= fnvPrime
	}
	return hash
}
