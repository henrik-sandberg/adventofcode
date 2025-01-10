package y2023

import (
	"adventofcode/shared"
)

func Day14(input []string) (solution shared.Solution[int, int]) {
	grid := shared.NewGrid(input)
	height, width := len(input), len(input[0])
	for w := 0; w < width; w++ {
		for h := 0; h < height; h++ {
			curr := complex(float32(w), float32(h))
			if grid[curr] != '.' {
				continue
			}
			for next := h + 1; next < height; next++ {
				n := complex(float32(w), float32(next))
				if grid[n] == '#' {
					break
				}
				if grid[n] == 'O' {
					grid[curr], grid[n] = grid[n], grid[curr]
					break
				}
			}
		}
	}
	for k, v := range grid {
		if v == 'O' {
			solution.Part1 += len(input) - int(imag(k))
		}
	}
	return
}
