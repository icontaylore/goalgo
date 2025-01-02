package main

import (
	"container/heap"
	"fmt"
)

type IntHeapMax []int

// реализац интерефейса heap
func (h IntHeapMax) Len() int {
	return len(h)
}
func (h IntHeapMax) Less(i int, j int) bool { return h[i] > h[j] }
func (h IntHeapMax) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }

// методы
func (h *IntHeapMax) Push(i interface{}) { *h = append(*h, i.(int)) }
func (h *IntHeapMax) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	h := &IntHeapMax{3, 4, 6, 7, 2, 50, 3}
	heap.Init(h)
	heap.Push(h, 56)
	fmt.Println(h)

}
