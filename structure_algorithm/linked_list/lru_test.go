package linkedlist

import (
	"fmt"
	"sync"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache(3)

	// lru.Put("a", 1)
	// lru.Put("b", 2)
	// lru.Put("c", 3)
	// lru.Put("d", 4)
	// lru.Put("e", 5)
	// lru.Get("a")
	// lru.Put("f", 6)
	// fmt.Println(lru.Get("f"))

	wg := sync.WaitGroup{}
	wg.Add(1)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			lru.Put(fmt.Sprintf("%d", i), i)
			// fmt.Println(i)
			wg.Done()
		}(i)

	}
	wg.Done()
	wg.Wait()
	fmt.Println("-----print------")
	lru.PrintlnAll()
}
