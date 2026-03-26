package lru_cache

import (
	"sync"

	"githhub.com/jrbarbati/lru-cache/internal/linked_list"
)

type LRUCache[T any] struct {
	capacity int
	mux      sync.Mutex
	cache    map[string]*linked_list.Node[T]
	usage    *linked_list.DoublyLinkedList[T]
	keyFunc  func(T) string
}

func New[T any](capacity int, keyFunc func(T) string) *LRUCache[T] {
	return &LRUCache[T]{
		capacity: capacity,
		mux:      sync.Mutex{},
		cache:    make(map[string]*linked_list.Node[T]),
		usage:    linked_list.New[T](),
		keyFunc:  keyFunc,
	}
}

func (lru *LRUCache[T]) Get(key string) (T, bool) {
	lru.mux.Lock()
	defer lru.mux.Unlock()

	if node, ok := lru.cache[key]; ok {
		lru.usage.PushFront(node)
		return node.Value, true
	}

	var zero T
	return zero, false
}

func (lru *LRUCache[T]) Put(value T) {
	lru.mux.Lock()
	defer lru.mux.Unlock()

	key := lru.keyFunc(value)

	if node, ok := lru.cache[key]; ok {
		node.Value = value
		lru.usage.MoveToFront(node)
		return
	}

	if len(lru.cache) >= lru.capacity {
		lru.removeLRU()
	}

	node := linked_list.NewNode(key, value)

	lru.cache[key] = node
	lru.usage.PushFront(node)
}

func (lru *LRUCache[T]) removeLRU() {
	back := lru.usage.Back()

	if back == nil {
		return
	}

	lru.usage.Remove(back)
	delete(lru.cache, lru.keyFunc(back.Value))
}
