package solutions

import "container/heap"

type MinHeap []int

// Pop implements heap.Interface.
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// Push implements heap.Interface.
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func KThLargestNumber(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}
	mh := &MinHeap{}
	heap.Init(mh)
	for _, a := range arr {
		heap.Push(mh, a)
		if len(*mh) > k {
			heap.Pop(mh)
		}
	}
	return (*mh)[0]
}
