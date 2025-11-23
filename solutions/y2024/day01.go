package y2024

import (
	"adventofcode/solutions/shared"
	"slices"
	"strings"
)

func Day01(input []string) (solution shared.Solution[int, int]) {
	var A []int
	var B []int
	for _, line := range input {
		ints := shared.IntSlice(strings.Fields(line))
		A = append(A, ints[0])
		B = append(B, ints[1])
	}
	slices.Sort(A)
	slices.Sort(B)
	for i := range input {
		solution.Part1 += shared.Abs(A[i] - B[i])
	}
	for _, v := range A {
		solution.Part2 += v * shared.Count(B, v)
	}
	return
}
