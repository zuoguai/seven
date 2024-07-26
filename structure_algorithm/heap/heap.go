package heap

type Heap []int

func NewHeap() *Heap {
	return &Heap{0}
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}
func (h *Heap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(v int) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *Heap) Pop() int {
	if len(*h) == 1 {
		return 0
	}
	v := (*h)[1]
	(*h)[1] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.down(1)
	return v
}

func (h Heap) up(u int) {
	t := u
	if u/2 > 0 && h.less(t, u/2) {
		t = u
	}
	if u != t {
		h.swap(u, t)
		h.up(t)
	}
}

func (h Heap) down(u int) {
	t := u
	if u*2 < len(h) && !h.less(t, u*2) {
		t = u * 2
	}
	if u*2+1 < len(h) && !h.less(t, u*2+1) {
		t = u*2 + 1
	}
	if u != t {
		h.swap(u, t)
		h.down(t)
	}

}
