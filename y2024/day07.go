package y2024

import (
	"adventofcode/shared"
	"strconv"
	"strings"
)

func Day07(input []string) (solution shared.Solution[int, int]) {
	for _, line := range input {
		solution.Part1 += validateEquation(line, []func(ints ...int) int{
			shared.Sum, shared.Product,
		})
		solution.Part2 += validateEquation(line, []func(ints ...int) int{
			shared.Sum, shared.Product,
			func(ints ...int) int {
				sb := strings.Builder{}
				for _, i := range ints {
					sb.WriteString(string(strconv.Itoa(i)))
				}
				newInt, _ := strconv.Atoi(sb.String())
				return newInt
			},
		})
	}
	return
}

func validateEquation(eq string, operations []func(ints ...int) int) int {
	equation := strings.Fields(eq)
	target, _ := strconv.Atoi(equation[0][:len(equation[0])-1])
	terms := shared.IntSlice(equation[1:])
	results := [][]int{terms}
	for len(results[0]) > 1 {
		nums := results[0]
		results = results[1:]
		for _, op := range operations {
			res := []int{op(nums[:2]...)}
			res = append(res, nums[2:]...)
			results = append(results, res)
		}
	}
	for _, res := range results {
		if res[0] == target {
			return target
		}
	}
	return 0
}
