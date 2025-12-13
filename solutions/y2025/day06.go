package y2025

import (
	"adventofcode/solutions/shared"
	"slices"
	"strconv"
	"strings"
)

func Day06(input []string) (solution shared.Solution[int, int]) {
	var nums [][]int
	for _, line := range input[:len(input)-1] {
		nums = append(nums, shared.IntSlice(strings.Fields(line)))
	}
	ops := input[len(input)-1]
	for idx, op := range strings.Fields(ops) {
		vals := slices.Collect(shared.Column(nums, idx))
		switch op {
		case "+":
			solution.Part1 += shared.Sum(vals...)
		case "*":
			solution.Part1 += shared.Product(vals...)
		}
	}
	var ns []int
	for idx := len(input[0]) - 1; idx >= 0; idx-- {
		var chars []byte
		for _, line := range input[:len(input)-1] {
			if line[idx] != ' ' {
				chars = append(chars, line[idx])
			}
		}
		if len(chars) > 0 {
			t, _ := strconv.Atoi(string(chars))
			ns = append(ns, t)
		}
		switch ops[idx] {
		case '+':
			solution.Part2 += shared.Sum(ns...)
			ns = nil
		case '*':
			solution.Part2 += shared.Product(ns...)
			ns = nil
		}
	}
	return
}
