package y2023

import (
	"adventofcode/solutions/shared"
	"container/heap"
)

func Day17(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	target := complex(float64(len(input[0])-1), float64(len(input)-1))
	solver := func(minSteps, maxSteps int) int {
		type position struct {
			tile      complex128
			direction complex128
		}
		pq := &shared.PriorityQueue[position]{}
		heap.Push(pq, &shared.Item[position]{Value: position{tile: 0, direction: 1}, Priority: 0})
		heap.Push(pq, &shared.Item[position]{Value: position{tile: 0, direction: 1i}, Priority: 0})
		seen := map[position]bool{}
		for len(*pq) > 0 {
			item := heap.Pop(pq).(*shared.Item[position])
			pos := item.Value
			if pos.tile == target {
				return item.Priority
			}
			if seen[pos] {
				continue
			}
			seen[pos] = true
			for _, turn := range []complex128{1i, -1i} {
				dir := pos.direction * turn
				for k := minSteps; k <= maxSteps; k++ {
					next := pos.tile + dir*complex(float64(k), 0)
					if _, ok := grid[next]; !ok {
						break
					}
					cost := item.Priority
					for step := 1; step <= k; step++ {
						cost += int(grid[pos.tile+dir*complex(float64(step), 0)] - '0')
					}
					heap.Push(pq, &shared.Item[position]{Value: position{next, dir}, Priority: cost})
				}
			}
		}
		panic("Did not reach target")
	}
	solution.Part1 = solver(1, 3)
	solution.Part2 = solver(4, 10)
	return
}
