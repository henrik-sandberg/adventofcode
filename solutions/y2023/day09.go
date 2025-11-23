package y2023

import (
	"adventofcode/shared"
	"strings"
)

func Day09(input []string) (solution shared.Solution[int, int]) {
	for _, line := range input {
		numbers := shared.IntSlice(strings.Fields(line))
		a, b := nextPolynomial(numbers)
		solution.Part1 += b
		solution.Part2 += a
	}
	return
}

func nextPolynomial(numbers []int) (int, int) {
	first := numbers[0]
	last := numbers[len(numbers)-1]
	if constantDiff(numbers) {
		return first, last
	}
	diffs := make([]int, len(numbers)-1)
	for i := range numbers[:len(numbers)-1] {
		diffs[i] = numbers[i+1] - numbers[i]
	}
	a, b := nextPolynomial(diffs)
	return first - a, last + b
}

func constantDiff(numbers []int) bool {
	for i := range numbers[:len(numbers)-2] {
		if numbers[i+2]-numbers[i+1] != numbers[i+1]-numbers[0] {
			return false
		}
	}
	return true
}
