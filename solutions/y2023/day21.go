package y2023

import (
	"adventofcode/solutions/shared"
	"iter"
	"maps"
)

func Day21(input []string) (solution shared.Solution[int, int]) {
	day21 := day21{input}
	solution.Part1 = day21.part1(64)
	solution.Part2 = day21.part2(26501365)
	return
}

type day21 struct {
	input []string
}

func (d *day21) bfs() iter.Seq[int] {
	grid := shared.NewGrid(d.input)
	v := grid.FindAny('S')
	grid[v] = '.'
	queue := []complex128{v}
	seen := map[complex128]int{v: 0}
	for len(queue) > 0 {
		v, queue = queue[0], queue[1:]
		for _, d := range []complex128{-1i, 1, 1i, -1} {
			next := v + d
			if _, ok := seen[next]; !ok && grid[next] == '.' {
				seen[next] = seen[v] + 1
				queue = append(queue, next)
			}
		}
	}
	return maps.Values(seen)
}

func (d *day21) part1(iterations int) (res int) {
	for v := range d.bfs() {
		if v <= iterations && v&1 == iterations&1 {
			res++
		}
	}
	return res
}

func (d *day21) part2(iterations int) int {
	n := iterations / len(d.input)
	evenFull := 0
	evenCorners := 0
	oddFull := 0
	oddCorners := 0
	for v := range d.bfs() {
		evenFull += v&1 ^ 1
		if v&1 == 0 && v > 65 {
			evenCorners++
		}
		oddFull += v & 1
		if v&1 == 1 && v > 65 {
			oddCorners++
		}
	}
	return (n+1)*(n+1)*oddFull + n*n*evenFull - (n+1)*oddCorners + n*evenCorners
}
