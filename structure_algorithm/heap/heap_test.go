package heap

import (
	"container/heap"
	"fmt"
	"math/rand"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap()
	res := []int{}
	data := []int{}
	for i := 0; i < 10; i++ {
		v := rand.Intn(88)
		data = append(data, v)
		h.Push(v)
	}
	fmt.Println(*h)
	for i := 0; i < 10; i++ {
		v := h.Pop()
		res = append(res, v)

	}
	fmt.Println(data)
	fmt.Println(res)

}

func TestHeap2(t *testing.T) {
	h := hp{}
	res := []int{}
	data := []int{}
	for i := 0; i < 10; i++ {
		v := rand.Intn(88)
		data = append(data, v)
		heap.Push(&h, v)
	}
	fmt.Println(h)
	for i := 0; i < 10; i++ {
		v := heap.Pop(&h)
		res = append(res, v.(int))

	}
	fmt.Println(data)
	fmt.Println(res)

}
