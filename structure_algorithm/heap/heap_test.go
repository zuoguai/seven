package heap

import (
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
