package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func kthLargest(arr []int, k int) int {
	h := &IntHeap{}
	for _, v := range arr {
		if h.Len() < k {
			heap.Push(h, v)
		} else if v > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}

	return (*h)[0]
}

func main() {
}
