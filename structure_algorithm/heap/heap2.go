package heap

type hp []int

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h hp) Len() int {
	return len(h)
}
func (h *hp) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}
func (h *hp) Push(val interface{}) {
	*h = append(*h, val.(int))
}
