package y2024

import (
	"adventofcode/shared"
)

func Day20(input []string) shared.Solution {
	return Day20Solver(input, 100)
}

func Day20Solver(input []string, limit int) shared.Solution {
	m := shared.NewGrid(input)
	start := m.FindAny('S')
	target := m.FindAny('E')
	arr := []complex64{start}
	for arr[len(arr)-1] != target {
		for _, dir := range []complex64{-1i, 1, 1i, -1} {
			next := arr[len(arr)-1] + dir
			if m[next] != '#' && (len(arr) < 2 || arr[len(arr)-2] != next) {
				arr = append(arr, next)
				break
			}
		}
	}
	part1 := 0
	part2 := 0
	manhattan := func(a, b complex64) int {
		return shared.Abs(int(real(a)-real(b))) + shared.Abs(int(imag(a)-imag(b)))
	}
	for i, first := range arr {
		for j := i + limit; j < len(arr); j++ {
			cheatDistance := manhattan(first, arr[j])
			saved := (j - i) - cheatDistance
			if cheatDistance <= 2 && saved >= limit {
				part1++
			}
			if cheatDistance <= 20 && saved >= limit {
				part2++
			}
		}
	}
	return shared.Solution{Part1: part1, Part2: part2}
}
