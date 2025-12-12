package y2025

import (
	"adventofcode/solutions/shared"
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
		tmp := nums[0][idx]
		for _, line := range nums[1:] {
			switch op {
			case "+":
				tmp += line[idx]
			case "*":
				tmp *= line[idx]
			}
		}
		solution.Part1 += tmp
	}
	var ns []int
	for idx := len(input[0]) - 1; idx >= 0; idx-- {
		var tmp []byte
		for _, line := range input[:len(input)-1] {
			if line[idx] != ' ' {
				tmp = append(tmp, line[idx])
			}
		}
		if t, _ := strconv.Atoi(string(tmp)); t != 0 {
			ns = append(ns, t)
		}
		if ops[idx] != ' ' {
			switch ops[idx] {
			case '+':
				solution.Part2 += shared.Sum(ns...)
			case '*':
				solution.Part2 += shared.Product(ns...)
			}
			ns = nil
		}
	}
	return
}
