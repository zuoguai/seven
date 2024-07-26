package linkedlist

import (
	"sync"
)

type Node struct {
	Next  *Node
	Prev  *Node
	Key   string
	Value interface{}
}

type LRUCache struct {
	head     *Node
	tail     *Node
	capacity int
	size     int
	cache    map[string]*Node
	sync.Mutex
}

func NewLRUCache(capacity int) *LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next, head.Prev = tail, tail
	tail.Next, tail.Prev = head, head

	return &LRUCache{
		head:     head,
		tail:     tail,
		capacity: capacity,
		cache:    make(map[string]*Node),
	}
}

func (lru *LRUCache) Len() int {
	return lru.size
}

func (lru *LRUCache) Get(key string) interface{} {
	lru.Lock()
	defer lru.Unlock()

	if lru.size == 0 {
		return nil
	}
	if lru.cache[key] != nil {
		lru.insertToHead(lru.cache[key])
		return lru.cache[key].Value
	}

	return nil
}

func (lru *LRUCache) Put(key string, value interface{}) {
	lru.Lock()
	defer lru.Unlock()
	if lru.cache[key] != nil {
		lru.cache[key].Value = value
		lru.insertToHead(lru.cache[key])
		return
	}
	if lru.size == lru.capacity {
		lru.deleteTail()
	}
	node := &Node{Key: key, Value: value}
	node.Next, node.Prev = lru.head.Next, lru.head
	lru.head.Next.Prev, lru.head.Next = node, node
	lru.cache[key] = node

	lru.size++
}

func (lru *LRUCache) PrintlnAll() {
	lru.Lock()
	defer lru.Unlock()
	for now := lru.head.Next; now != lru.tail; now = now.Next {
		println(now.Key)
	}
}

func (lru *LRUCache) insertToHead(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
	node.Next, node.Prev = lru.head.Next, lru.head
	lru.head.Next, lru.head.Next.Prev = node, node
}

func (lru *LRUCache) deleteTail() {
	delete(lru.cache, lru.tail.Prev.Key)
	lru.tail.Prev.Prev.Next = lru.tail
	lru.tail.Prev = lru.tail.Prev.Prev
	lru.size--
}
