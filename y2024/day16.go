package y2024

import (
	"adventofcode/shared"
	"container/heap"
	"math"
)

func Day16(input []string) (solution shared.Solution[int, int]) {
	m := shared.NewGrid(input)
	type pathkey struct {
		position  complex128
		direction complex128
	}
	start := m.FindAny('S')
	target := m.FindAny('E')
	m[start], m[target] = '.', '.'
	best := math.MaxInt
	dist := map[pathkey]int{{start, 1}: 0}
	seen := map[complex128]bool{start: true}
	pq := &shared.PriorityQueue[reindeer]{}
	heap.Init(pq)
	enqueue := func(r reindeer) {
		heap.Push(pq, &shared.Item[reindeer]{Value: r, Priority: r.cost})
	}
	enqueue(reindeer{start, 1, 0, map[complex128]bool{start: true}})
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*shared.Item[reindeer]).Value
		key := pathkey{current.position, current.direction}
		if previous, ok := dist[key]; ok && current.cost > previous {
			continue
		}
		dist[key] = current.cost
		if current.position == target && current.cost <= best {
			for k := range current.seen {
				seen[k] = true
			}
			best = current.cost
		}
		for _, change := range []struct {
			dir  complex128
			cost int
		}{{1, 1}, {-1i, 1001}, {1i, 1001}} {
			dir := current.direction * change.dir
			next := current.position + dir
			key := pathkey{next, dir}
			if _, ok := dist[key]; ok {
				continue
			}
			if m[next] == '.' {
				p := make(map[complex128]bool, len(current.seen)+1)
				for k, v := range current.seen {
					p[k] = v
				}
				p[next] = true
				enqueue(reindeer{next, dir, current.cost + change.cost, p})
			}
		}
	}
	solution.Part1 = best
	solution.Part2 = len(seen)
	return
}

type reindeer struct {
	position  complex128
	direction complex128
	cost      int
	seen      map[complex128]bool
}
