package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)

	var arrLen int
	if _, err := fmt.Fscan(reader, &arrLen); err != nil {
		return
	}

	arr := make([]int, arrLen)
	for i := range arr {
		if _, err := fmt.Fscan(reader, &arr[i]); err != nil {
			return
		}
	}

	var k int
	if _, err := fmt.Fscan(reader, &k); err != nil {
		return
	}

	if k < 1 || k > arrLen {
		_, _ = fmt.Println(-1)

		return
	}

	_, _ = fmt.Println(kthLargest(arr, k))
}
