package main

import (
	"container/heap"
	"fmt"

	"github.com/Madmaxvas/task-2-2/maxheap"
)

func main() {
	var dishesCount int
	if _, err := fmt.Scan(&dishesCount); err != nil {
		fmt.Println(fmt.Errorf("ошибка чтения количества блюд: %w", err))

		return
	}

	maxHeap := &maxheap.IntHeap{}

	for range dishesCount {
		var val int
		if _, err := fmt.Scan(&val); err != nil {
			fmt.Println(fmt.Errorf("ошибка чтения оценки блюда: %w", err))

			return
		}

		*maxHeap = append(*maxHeap, val)
	}

	heap.Init(maxHeap)

	var targetRank int
	if _, err := fmt.Scan(&targetRank); err != nil {
		fmt.Println(fmt.Errorf("ошибка чтения искомой позиции: %w", err))

		return
	}

	var result int

	for range targetRank {
		popped := heap.Pop(maxHeap)
		if popped == nil {
			fmt.Println("ошибка: куча пуста")

			return
		}

		val, ok := popped.(int)
		if !ok {
			fmt.Println("ошибка: неверный тип извлеченного элемента")

			return
		}

		result = val
	}

	fmt.Println(result)
}
