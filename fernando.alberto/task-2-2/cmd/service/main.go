package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func kthLargest(arr []int, kValue int) int {
	minHeap := &IntHeap{}
	for _, v := range arr {
		if minHeap.Len() < kValue {
			heap.Push(minHeap, v)
		} else if v > (*minHeap)[0] {
			heap.Pop(minHeap)
			heap.Push(minHeap, v)
		}
	}

	return (*minHeap)[0]
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

	var kValue int
	if _, err := fmt.Fscan(reader, &kValue); err != nil {
		return
	}

	if kValue < 1 || kValue > arrLen {
		_, _ = fmt.Println(-1)

		return
	}

	_, _ = fmt.Println(kthLargest(arr, kValue))
}
