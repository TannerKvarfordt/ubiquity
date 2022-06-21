package datastructures

import (
	"bytes"
	"encoding/gob"
	"sync"
)

const ShardCount = 128

// ConcurrentMap is a sharded, thread-safe mapping of keys to values.
type ConcurrentMap[KEY comparable, VAL any] struct {
	shards [ShardCount]*shard[KEY, VAL]
}

// NewConcurrentMap creates a new ConcurrentMap.
func NewConcurrentMap[KEY comparable, VAL any]() *ConcurrentMap[KEY, VAL] {
	c := ConcurrentMap[KEY, VAL]{
		shards: [ShardCount]*shard[KEY, VAL]{},
	}
	for i := range c.shards {
		c.shards[i] = newShard[KEY, VAL]()
	}
	return &c
}

func (c *ConcurrentMap[KEY, VAL]) getShard(key KEY) *shard[KEY, VAL] {
	return c.shards[fnv1a(key)%uint64(ShardCount)]
}

func (c *ConcurrentMap[KEY, VAL]) Reset() {
	for _, shard := range c.shards {
		shard.mutex.Lock()
		for key := range shard.Contents {
			delete(shard.Contents, key)
		}
		shard.mutex.Unlock()
	}
}

// At returns the value for the given key and a bool indicating whether or not
// the given key exists in the map. If the given key does not exist, then the
// zero-value of VAL and false are returned.
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

// SetIfAbsent inserts the given val at the specified key, so long as
// the specified key does not already exist in the map. A bool is returned
// indicating whether or not the insert took place.
func (c *ConcurrentMap[KEY, VAL]) SetIfAbsent(key KEY, val VAL) bool {
	shard := c.getShard(key)
	return shard.setIfAbsent(key, val)
}

// SetIfPresent inserts the given val at the specified key, so long as
// the specified key already exists in the map. A bool is returned
// indicating whether or not the insert took place.
func (c *ConcurrentMap[KEY, VAL]) SetIfPresent(key KEY, val VAL) bool {
	shard := c.getShard(key)
	return shard.setIfPresent(key, val)
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

// Data returns the contents of all ConcurrentMap shards as a map[KEY]VAL.
func (c *ConcurrentMap[KEY, VAL]) Data() map[KEY]VAL {
	m := make(map[KEY]VAL, c.Len())
	for _, shard := range c.shards {
		shard.mutex.RLock()
		for key, val := range shard.Contents {
			m[key] = val
		}
		shard.mutex.RUnlock()
	}
	return m
}

// Remove deletes the element with the specified key from the
// map, if it exists. A boolean is returned indicating whether
// or not any deletion was necessary.
func (c *ConcurrentMap[KEY, VAL]) Remove(key KEY) bool {
	shard := c.getShard(key)
	return shard.remove(key)
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

func (s *shard[KEY, VAL]) setIfAbsent(k KEY, v VAL) bool {
	if _, ok := s.at(k); !ok {
		s.set(k, v)
		return true
	}
	return false
}

func (s *shard[KEY, VAL]) setIfPresent(k KEY, v VAL) bool {
	if _, ok := s.at(k); !ok {
		return false
	}
	s.set(k, v)
	return true
}

func (s *shard[KEY, VAL]) remove(k KEY) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, ok := s.Contents[k]
	delete(s.Contents, k)
	return ok
}
