package shared

type PriorityQueue[T any] []*Item[T]

type Item[T any] struct {
	Value    T
	Priority int // Lower is higher priority
}

func (pq *PriorityQueue[T]) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Priority < (*pq)[j].Priority
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
