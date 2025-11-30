package y2021

import (
	"adventofcode/solutions/shared"
	"math"
	"slices"
	"strings"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	type board struct {
		values []int
		marks  []bool
	}
	isSolved := func(b board) bool {
		for _, winningSquares := range [][]int{
			{0, 1, 2, 3, 4},
			{5, 6, 7, 8, 9},
			{10, 11, 12, 13, 14},
			{15, 16, 17, 18, 19},
			{20, 21, 22, 23, 24},
			{0, 5, 10, 15, 20},
			{1, 6, 11, 16, 21},
			{2, 7, 12, 17, 22},
			{3, 8, 13, 18, 23},
			{4, 9, 14, 19, 24},
		} {
			solved := true
			for _, idx := range winningSquares {
				if !b.marks[idx] {
					solved = false
					break
				}
			}
			if solved {
				return true
			}
		}
		return false
	}
	moves := shared.IntSlice(strings.Split(input[0], ","))
	first := math.MaxInt
	second := 0
	for row := 2; row < len(input); row += 6 {
		b := board{
			values: shared.IntSlice(strings.Fields(strings.Join(input[row:row+5], " "))),
			marks:  make([]bool, 25),
		}
		for used, move := range moves {
			idx := slices.Index(b.values, move)
			if idx == -1 {
				continue
			}
			b.marks[idx] = true
			if !isSolved(b) {
				continue
			}
			unmarked := 0
			for i := range 25 {
				if !b.marks[i] {
					unmarked += b.values[i]
				}
			}
			if used < first {
				first = used
				solution.Part1 = unmarked * move
			}
			if used > second {
				second = used
				solution.Part2 = unmarked * move
			}
			break
		}
	}
	return
}
