package y2022

import (
	"adventofcode/solutions/shared"
	"strconv"
	"strings"
)

func Day14(input []string) (solution shared.Solution[int, int]) {
	start := complex128(500)
	maxY := 0.0
	grid := make(map[complex128]bool)
	parse := func(s string) (int, int) {
		x, y, _ := strings.Cut(s, ",")
		xi, _ := strconv.Atoi(x)
		yi, _ := strconv.Atoi(y)
		return xi, yi
	}
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		for i := range len(parts) - 1 {
			x0, y0 := parse(parts[i])
			x1, y1 := parse(parts[i+1])
			dx := shared.Sign(x1 - x0)
			dy := shared.Sign(y1 - y0)
			steps := max(shared.Abs(x1-x0), shared.Abs(y1-y0))
			for s := 0; s <= steps; s++ {
				x := float64(x0 + s*dx)
				y := float64(y0 + s*dy)
				if y > maxY {
					maxY = y
				}
				grid[complex(x, y)] = true
			}
		}
	}
	drop := func(p complex128) bool {
		if grid[p] {
			return false
		}
		for {
			if imag(p) > maxY {
				return false
			}
			down := p + 1i
			if !grid[down] {
				p = down
				continue
			}
			left := p - 1 + 1i
			if !grid[left] {
				p = left
				continue
			}
			right := p + 1 + 1i
			if !grid[right] {
				p = right
				continue
			}
			grid[p] = true
			return true
		}
	}
	for drop(start) {
		solution.Part1++
		solution.Part2++
	}
	maxY += 2
	for i := real(start) - maxY; i <= real(start)+maxY; i++ {
		grid[complex(i, maxY)] = true
	}
	for drop(start) {
		solution.Part2++
	}
	return
}
