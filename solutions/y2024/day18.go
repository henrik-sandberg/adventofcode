package y2024

import (
	"strings"

	"adventofcode/solutions/shared"
)

func Day18(input []string) (solution shared.Solution[int, string]) {
	iterations := 1024
	size := 71
	target := complex128(70 + 70i)
	m := map[complex128]rune{}
	for r := range size {
		for c := range size {
			m[complex(float64(c), float64(r))] = '.'
		}
	}
	dropByte := func(point string) {
		i := shared.IntSlice(strings.Split(point, ","))
		m[complex(float64(i[0]), float64(i[1]))] = '#'
	}
	bfs := func() int {
		current := complex128(0)
		queue := []complex128{current}
		seen := map[complex128]int{current: 0}
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]
			for _, dir := range []complex128{-1i, 1, 1i, -1} {
				next := current + dir
				if next == target {
					return seen[current] + 1
				}
				_, visited := seen[next]
				if !visited && m[next] == '.' {
					queue = append(queue, next)
					seen[next] = seen[current] + 1
				}
			}
		}
		return -1
	}
	for _, line := range input[:iterations] {
		dropByte(line)
	}
	solution.Part1 = bfs()
	for _, line := range input[iterations:] {
		dropByte(line)
		if bfs() == -1 {
			solution.Part2 = line
			break
		}
	}
	return
}
