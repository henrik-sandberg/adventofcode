package y2021

import (
	"adventofcode/solutions/shared"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	isSolved := func(b []string) bool {
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
				if b[idx] != "x" {
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
	moves := strings.Split(input[0], ",")
	first := math.MaxInt
	second := 0
	for i := 2; i < len(input); i += 6 {
		board := strings.Fields(strings.Join(input[i:i+5], " "))
		for used, m := range moves {
			ind := slices.Index(board, m)
			if ind == -1 {
				continue
			}
			board[ind] = "x"
			if !isSolved(board) {
				continue
			}
			last, _ := strconv.Atoi(m)
			if used < first {
				first = used
				solution.Part1 = shared.Sum(shared.IntSlice(board)...) * last
			}
			if used > second {
				second = used
				solution.Part2 = shared.Sum(shared.IntSlice(board)...) * last
			}
			break
		}
	}
	return
}
