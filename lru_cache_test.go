package lru_cache

import (
	"strconv"
	"sync"
	"testing"

	"githhub.com/jrbarbati/lru-cache/internal/linked_list"
)

var (
	getScenarios = []struct {
		name          string
		prepareCache  func() *LRUCache[int]
		getKey        string
		expectValue   bool
		expectedValue int
	}{
		{
			name: "Empty Cache",
			prepareCache: func() *LRUCache[int] {
				return &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache:    map[string]*linked_list.Node[int]{},
					usage:    linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}
			},
			getKey:        "0",
			expectValue:   false,
			expectedValue: 0,
		},
		{
			name: "Non Empty Cache Hit",
			prepareCache: func() *LRUCache[int] {
				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": linked_list.NewNode[int]("1", 1),
					},
					usage: linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}

				for _, node := range lruCache.cache {
					lruCache.usage.PushFront(node)
				}

				return lruCache
			},
			getKey:        "1",
			expectValue:   true,
			expectedValue: 1,
		},
		{
			name: "Non Empty Cache Miss",
			prepareCache: func() *LRUCache[int] {
				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": linked_list.NewNode[int]("1", 1),
					},
					usage: linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}

				for _, node := range lruCache.cache {
					lruCache.usage.PushFront(node)
				}

				return lruCache
			},
			getKey:        "2",
			expectValue:   false,
			expectedValue: 0,
		},
	}
	putScenarios = []struct {
		name          string
		startingCache func() *LRUCache[int]
		value         int
		expectedCache func() *LRUCache[int]
	}{
		{
			name: "Empty Cache",
			startingCache: func() *LRUCache[int] {
				return &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache:    map[string]*linked_list.Node[int]{},
					usage:    linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}
			},
			value: 1,
			expectedCache: func() *LRUCache[int] {
				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": linked_list.NewNode[int]("1", 1),
					},
					usage: linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}

				for _, node := range lruCache.cache {
					lruCache.usage.PushFront(node)
				}

				return lruCache
			},
		},
		{
			name: "Non Empty Cache",
			startingCache: func() *LRUCache[int] {
				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": linked_list.NewNode[int]("1", 1),
					},
					usage: linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}

				for _, node := range lruCache.cache {
					lruCache.usage.PushFront(node)
				}

				return lruCache
			},
			value: 2,
			expectedCache: func() *LRUCache[int] {
				node1 := linked_list.NewNode[int]("1", 1)
				node2 := linked_list.NewNode[int]("2", 2)

				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": node1,
						"2": node2,
					},
					usage:   linked_list.New[int](),
					keyFunc: func(v int) string { return strconv.Itoa(v) },
				}

				lruCache.usage.PushFront(node2)
				lruCache.usage.PushFront(node1)
				return lruCache
			},
		},
		{
			name: "Non Empty Cache -- Already Exists",
			startingCache: func() *LRUCache[int] {
				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": linked_list.NewNode[int]("1", 1),
					},
					usage: linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}

				for _, node := range lruCache.cache {
					lruCache.usage.PushFront(node)
				}

				return lruCache
			},
			value: 1,
			expectedCache: func() *LRUCache[int] {
				node1 := linked_list.NewNode[int]("1", 1)

				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": node1,
					},
					usage:   linked_list.New[int](),
					keyFunc: func(v int) string { return strconv.Itoa(v) },
				}

				lruCache.usage.PushFront(node1)
				return lruCache
			},
		},
		{
			name: "Full Cache",
			startingCache: func() *LRUCache[int] {
				node1 := linked_list.NewNode[int]("1", 1)
				node2 := linked_list.NewNode[int]("2", 2)

				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": node1,
						"2": node2,
					},
					usage:   linked_list.New[int](),
					keyFunc: func(v int) string { return strconv.Itoa(v) },
				}

				lruCache.usage.PushFront(node2)
				lruCache.usage.PushFront(node1)
				return lruCache
			},
			value: 3,
			expectedCache: func() *LRUCache[int] {
				node1 := linked_list.NewNode[int]("1", 1)
				node3 := linked_list.NewNode[int]("3", 3)

				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": node1,
						"3": node3,
					},
					usage:   linked_list.New[int](),
					keyFunc: func(v int) string { return strconv.Itoa(v) },
				}

				lruCache.usage.PushFront(node3)
				lruCache.usage.PushFront(node1)
				return lruCache
			},
		},
	}
	removeScenarios = []struct {
		name          string
		startingCache func() *LRUCache[int]
		expectedCache func() *LRUCache[int]
	}{
		{
			name: "Full Cache",
			startingCache: func() *LRUCache[int] {
				node1 := linked_list.NewNode[int]("1", 1)
				node2 := linked_list.NewNode[int]("2", 2)

				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": node1,
						"2": node2,
					},
					usage:   linked_list.New[int](),
					keyFunc: func(v int) string { return strconv.Itoa(v) },
				}

				lruCache.usage.PushFront(node2)
				lruCache.usage.PushFront(node1)
				return lruCache
			},
			expectedCache: func() *LRUCache[int] {
				lruCache := &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache: map[string]*linked_list.Node[int]{
						"1": linked_list.NewNode[int]("1", 1),
					},
					usage: linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}

				for _, node := range lruCache.cache {
					lruCache.usage.PushFront(node)
				}

				return lruCache
			},
		},
		{
			name: "Empty Cache",
			startingCache: func() *LRUCache[int] {
				return &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache:    map[string]*linked_list.Node[int]{},
					usage:    linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}
			},
			expectedCache: func() *LRUCache[int] {
				return &LRUCache[int]{
					capacity: 2,
					mux:      sync.Mutex{},
					cache:    map[string]*linked_list.Node[int]{},
					usage:    linked_list.New[int](),
					keyFunc: func(v int) string {
						return strconv.Itoa(v)
					},
				}
			},
		},
	}
)

func TestLRUCache_Get(t *testing.T) {
	for _, scenario := range getScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			value, ok := scenario.prepareCache().Get(scenario.getKey)

			if scenario.expectValue && !ok {
				t.Fatalf("Expected value to be %t, no value found", scenario.expectValue)
			}

			if !scenario.expectValue && ok {
				t.Fatalf("Expected no value, but got one %v", value)
			}

			if value != scenario.expectedValue {
				t.Fatalf("Expected value to be %v, got %v", scenario.expectedValue, value)
			}
		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	for _, scenario := range putScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			lruCache := scenario.startingCache()
			lruCache.Put(scenario.value)

			if len(scenario.expectedCache().cache) != len(lruCache.cache) {
				t.Fatalf("Cache is not expected size (%v): %v", len(scenario.expectedCache().cache), len(lruCache.cache))
			}

			for key, expectedValue := range scenario.expectedCache().cache {
				value, ok := lruCache.Get(key)

				if !ok {
					t.Fatalf("Expected key not found in cache: %v", key)
				}

				if expectedValue.Value != value {
					t.Fatalf("Expected value to be %v, got %v", expectedValue, value)
				}
			}
		})
	}
}

func TestLRUCache_removeLRUEntry(t *testing.T) {
	for _, scenario := range removeScenarios {
		t.Run(scenario.name, func(t *testing.T) {
			lruCache := scenario.startingCache()
			lruCache.removeLRU()

			if len(scenario.expectedCache().cache) != len(lruCache.cache) {
				t.Fatalf("Cache is not expected size (%v): %v", len(scenario.expectedCache().cache), len(lruCache.cache))
			}

			for key, expectedValue := range scenario.expectedCache().cache {
				value, ok := lruCache.Get(key)

				if !ok {
					t.Fatalf("Expected key not found in cache: %v", key)
				}

				if expectedValue.Value != value {
					t.Fatalf("Expected value to be %v, got %v", expectedValue, value)
				}
			}
		})
	}
}

func BenchmarkLRUCache_Get(b *testing.B) {
	b.ReportAllocs()
	for _, scenario := range getScenarios {
		b.Run(scenario.name, func(b *testing.B) {
			lruCache := scenario.prepareCache()
			b.ResetTimer()

			for b.Loop() {
				lruCache.Get(scenario.getKey)
			}
		})
	}
}

func BenchmarkLRUCache_Put(b *testing.B) {
	b.ReportAllocs()
	for _, scenario := range putScenarios {
		b.Run(scenario.name, func(b *testing.B) {
			lruCache := scenario.startingCache()
			b.ResetTimer()

			for b.Loop() {
				lruCache.Put(scenario.value)
			}
		})
	}
}
