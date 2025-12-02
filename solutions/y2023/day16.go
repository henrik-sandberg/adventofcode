package y2023

import (
	"adventofcode/solutions/shared"
)

func Day16(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	type light struct {
		pos, dir complex128
	}
	reflect := func(l light) []complex128 {
		switch grid[l.pos] {
		case '|':
			if real(l.dir) != 0 {
				return []complex128{-1i, 1i}
			}
		case '-':
			if imag(l.dir) != 0 {
				return []complex128{-1, 1}
			}
		case '/':
			if real(l.dir) != 0 {
				return []complex128{l.dir * -1i}
			}
			return []complex128{l.dir * 1i}
		case '\\':
			if real(l.dir) != 0 {
				return []complex128{l.dir * 1i}
			}
			return []complex128{l.dir * -1i}
		}
		return []complex128{l.dir}
	}
	solver := func(l light) int {
		lights := []light{l}
		seen := map[light]bool{l: true}
		for len(lights) > 0 {
			l, lights = lights[0], lights[1:]
			for _, dir := range reflect(l) {
				next := light{l.pos + dir, dir}
				if _, ok := grid[next.pos]; ok && !seen[next] {
					seen[next] = true
					lights = append(lights, next)
				}
			}
		}
		unique := map[complex128]any{}
		for k := range seen {
			unique[k.pos] = nil
		}
		return len(unique)
	}
	solution.Part1 = solver(light{0, 1})
	c := func(a, b int) complex128 {
		return complex(float64(a), float64(b))
	}
	for i := range len(input) {
		solution.Part2 = max(solution.Part2, solver(light{c(0, i), 1}))
		solution.Part2 = max(solution.Part2, solver(light{c(i, 0), 1i}))
		solution.Part2 = max(solution.Part2, solver(light{c(i, len(input)-1), -1i}))
		solution.Part2 = max(solution.Part2, solver(light{c(len(input[i])-1, i), -1}))
	}
	return
}
