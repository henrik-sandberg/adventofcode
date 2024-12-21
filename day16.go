package main

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

func Day16(input []string) {
	m := Map{strings.Join(input, ""), len(input[0])}.ToComplexGrid()
	type pathkey struct {
		position  complex64
		direction complex64
	}
	var start, target complex64
	for k, v := range m {
		switch v {
		case 'S':
			start = k
			m[k] = '.'
		case 'E':
			target = k
			m[k] = '.'
		}
	}
	fmt.Println("Start:", start, "target:", target)
	best := math.MaxInt
	dist := map[pathkey]int{{start, 1}: 0}
	seen := map[complex64]bool{start: true}
	pq := &PriorityQueue[reindeer]{}
	heap.Init(pq)
	enqueue := func(r reindeer) {
		heap.Push(pq, &Item[reindeer]{value: r, priority: r.cost})
	}
	enqueue(reindeer{start, 1, 0, map[complex64]bool{start: true}})
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item[reindeer]).value
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
			dir  complex64
			cost int
		}{{1, 1}, {-1i, 1001}, {1i, 1001}} {
			dir := current.direction * change.dir
			next := current.position + dir
			key := pathkey{next, dir}
			if _, ok := dist[key]; ok {
				continue
			}
			if m[next] == '.' {
				p := make(map[complex64]bool, len(current.seen)+1)
				for k, v := range current.seen {
					p[k] = v
				}
				p[next] = true
				enqueue(reindeer{next, dir, current.cost + change.cost, p})
			}
		}
	}
	fmt.Println(best)
	fmt.Println(len(seen))
}

type reindeer struct {
	position  complex64
	direction complex64
	cost      int
	seen      map[complex64]bool
}
