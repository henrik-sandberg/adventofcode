package y2025

import (
	"strconv"

	"adventofcode/solutions/shared"
)

func Day01(input []string) (solution shared.Solution[int, int]) {
	current := 50
	for _, step := range input {
		val, _ := strconv.Atoi(step[1:])

		solution.Part2 += val / 100
		rem := val % 100

		if step[0] == 'L' {
			if current > 0 && rem >= current {
				solution.Part2++
			}
			current -= rem
		} else {
			if current+rem >= 100 {
				solution.Part2++
			}
			current += rem
		}

		current = shared.PositiveMod(current, 100)
		if current == 0 {
			solution.Part1++
		}
	}
	return
}
