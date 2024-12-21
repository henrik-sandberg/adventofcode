package main

type Item[T any] struct {
	value    T
	priority int // Lower is higher priority
}

type PriorityQueue[T any] []*Item[T]

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].priority < (*pq)[j].priority
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	*pq = append(*pq, x.(*Item[T]))

}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
