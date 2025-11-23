package y2023

import (
	"adventofcode/solutions/shared"
	"slices"
)

func Day23(input []string) (solution shared.Solution[int, int]) {
	solver := day23{
		grid:   shared.NewGrid(input),
		target: complex(float64(len(input[0])-2), float64(len(input)-1)),
		start:  complex(1, 0),
	}
	solution.Part1 = solver.part1()
	solution.Part2 = solver.part2()
	return
}

type day23 struct {
	grid          shared.Grid
	start, target complex128
}

func (d *day23) part1() int {
	return d.dfs([]complex128{d.start})
}

func (d *day23) part2() int {
	type tile struct {
		pos  complex128
		cost int
	}
	intersections := d.intersections()
	graph := map[complex128][]tile{}
graphbuilder:
	for c := range shared.Combinations(intersections, 2) {
		from, target := c[0], c[1]
		queue := []tile{{from, 0}}
		seen := map[complex128]bool{}
		var t tile
		for len(queue) > 0 {
			t, queue = queue[0], queue[1:]
			for _, dir := range d.grid.Directions() {
				next := t.pos + dir
				if next == target {
					graph[from] = append(graph[from], tile{target, t.cost + 1})
					graph[target] = append(graph[target], tile{from, t.cost + 1})
					continue graphbuilder
				}
				val, ok := d.grid[next]
				if ok && val != '#' && !seen[next] && !slices.Contains(intersections, next) {
					seen[next] = true
					queue = append(queue, tile{next, t.cost + 1})
				}
			}
		}
	}
	var dfs func([]tile) int
	dfs = func(path []tile) (cost int) {
		for _, p := range graph[path[len(path)-1].pos] {
			if p.pos == d.target {
				cost = p.cost
				for _, t := range path {
					cost += t.cost
				}
				return
			}
			if !slices.ContainsFunc(path, func(t tile) bool {
				return t.pos == p.pos
			}) {
				path = append(path, p)
				cost = max(cost, dfs(path))
				path = path[:len(path)-1]
			}
		}
		return
	}
	return dfs([]tile{{d.start, 0}})
}

func (d *day23) dfs(path []complex128) (cost int) {
	directions := map[byte]complex128{'>': 1, 'v': 1i, '<': -1, '^': -1i}
	for _, dir := range d.grid.Directions() {
		next := path[len(path)-1] + dir
		if next == d.target {
			return len(path)
		}
		val := d.grid[next]
		if val != '#' && (val == '.' || dir == directions[val]) && !slices.Contains(path, next) {
			path = append(path, next)
			cost = max(cost, d.dfs(path))
			path = path[:len(path)-1]
		}
	}
	return
}

func (d *day23) intersections() []complex128 {
	intersections := []complex128{d.start, d.target}
	for v, val := range d.grid {
		if val != '#' {
			count := 0
			for _, dir := range d.grid.Directions() {
				val, ok := d.grid[v+dir]
				if ok && val != '#' {
					count++
				}
			}
			if count > 2 {
				intersections = append(intersections, v)
			}
		}
	}
	return intersections
}
