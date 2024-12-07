package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day07(input []string) {
	part1 := 0
	part2 := 0
	for _, line := range input {
		part1 += validateEquation(line, []func(ints ...int) int{
			sum, multiply,
		})
		part2 += validateEquation(line, []func(ints ...int) int{
			sum, multiply,
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
	fmt.Println(part1)
	fmt.Println(part2)
}

func validateEquation(eq string, operations []func(ints ...int) int) int {
	equation := strings.Fields(eq)
	target, _ := strconv.Atoi(equation[0][:len(equation[0])-1])
	terms := ToInts(equation[1:])
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
