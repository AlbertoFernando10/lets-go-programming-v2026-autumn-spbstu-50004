package maxheap

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	if i < 0 || j < 0 || i >= len(*h) || j >= len(*h) {
		return false
	}

	return (*h)[i] > (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	if i < 0 || j < 0 || i >= len(*h) || j >= len(*h) {
		return
	}

	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x any) {
	val, ok := x.(int)
	if !ok {
		return
	}

	*h = append(*h, val)
}

func (h *IntHeap) Pop() any {
	old := *h
	oldLen := len(old)

	if oldLen == 0 {
		return nil
	}

	x := old[oldLen-1]
	*h = old[0 : oldLen-1]

	return x
}
