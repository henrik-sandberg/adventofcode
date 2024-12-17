package main

import (
	"fmt"
	"math"
	"strings"
)

func Day16(input []string) {
	m := Map{strings.Join(input, ""), len(input[0])}.ToComplexGrid()
	type path struct {
		cost  int
		tiles map[complex64]bool
	}
	type pathkey struct {
		position  complex64
		direction complex64
	}
	var start complex64
	for k, v := range m {
		if v == 'S' {
			start = k
		}
	}
	current := reindeer{start, 1, 0, map[complex64]bool{start: true}}
	paths := map[pathkey]path{}
	queue := []reindeer{current}
	failfast := map[complex64]int{}
	best := math.MaxInt
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current.cost > best {
			continue
		} else if m[current.position] == 'E' && current.cost < best {
			best = current.cost
		}
		if seen, ok := failfast[current.position]; !ok || current.cost < seen+1001 {
			failfast[current.position] = current.cost
		} else {
			continue
		}
		key := pathkey{current.position, current.direction}
		previous, ok := paths[key]
		if ok && current.cost > previous.cost {
			continue
		}
		if !ok || current.cost < previous.cost {
			paths[key] = path{current.cost, current.seen}
		} else if current.cost == previous.cost {
			for k, v := range previous.tiles {
				current.seen[k] = v
			}
			paths[key] = path{current.cost, current.seen}
		}
		if v := m[current.position+current.direction]; v != '#' {
			queue = append(queue, current.move(current.direction))
		}
		if left := current.direction * 1i; m[current.position+left] != '#' {
			queue = append(queue, current.move(left))
		}
		if right := current.direction * -1i; m[current.position+right] != '#' {
			queue = append(queue, current.move(right))
		}
	}

	for k, v := range paths {
		if m[k.position] == 'E' {
			fmt.Println(k, v.cost, len(v.tiles))
		}
	}
}

type reindeer struct {
	position  complex64
	direction complex64
	cost      int
	seen      map[complex64]bool
}

func (r reindeer) move(direction complex64) reindeer {
	cost := r.cost
	nextPosition := r.position
	if r.direction == direction {
		cost += 1
		nextPosition += direction
	} else {
		cost += 1000
	}
	p := make(map[complex64]bool, len(r.seen)+1)
	for k, v := range r.seen {
		p[k] = v
	}
	p[nextPosition] = true
	return reindeer{nextPosition, direction, cost, p}
}
