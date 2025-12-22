package y2025

import (
	"adventofcode/solutions/shared"
	"strconv"
	"strings"
)

func Day10(input []string) (solution shared.Solution[int, int]) {
	encode := func(a []int) string {
		parts := make([]string, len(a))
		for i, v := range a {
			parts[i] = strconv.Itoa(v)
		}
		return strings.Join(parts, ",")
	}
	type Pattern struct {
		value []int
		cost  int
	}
	makeCoeffs := func(buttons [][]int, targetLength int) [][]int {
		coeffs := make([][]int, len(buttons))
		for idx := range buttons {
			c := make([]int, targetLength)
			for _, i := range buttons[idx] {
				c[i] = 1
			}
			coeffs[idx] = c
		}
		return coeffs
	}
	makePatterns := func(coeffs [][]int) []Pattern {
		indices := make([]int, len(coeffs))
		for i := range coeffs {
			indices[i] = i
		}
		numVariables := len(coeffs[0])
		seen := make(map[string]struct{})
		var out []Pattern
		for k := range len(coeffs) + 1 {
			for comb := range shared.Combinations(indices, k) {
				work := make([]int, numVariables)
				for _, i := range comb {
					for j := range numVariables {
						work[j] += coeffs[i][j]
					}
				}
				key := encode(work)
				if _, ok := seen[key]; !ok {
					seen[key] = struct{}{}
					out = append(out, Pattern{value: work, cost: k})
				}
			}
		}
		return out
	}
	findParity := func(patterns []Pattern, target []int) int {
		for _, pattern := range patterns {
			valid := true
			for idx := range target {
				if pattern.value[idx]&1 != target[idx]&1 {
					valid = false
					break
				}
			}
			if valid {
				return pattern.cost
			}
		}
		return -1
	}
	findAddends := func(patterns []Pattern, target []int) int {
		memo := map[string]int{
			encode(make([]int, len(target))): 0,
		}
		var helper func(target []int) int
		helper = func(target []int) int {
			key := encode(target)
			if v, ok := memo[key]; ok {
				return v
			}
			ans := 1 << 60
			for _, pattern := range patterns {
				valid := true
				for idx := range target {
					if pattern.value[idx] > target[idx] || pattern.value[idx]&1 != target[idx]&1 {
						valid = false
						break
					}
				}
				if !valid {
					continue
				}
				next := make([]int, len(target))
				for idx := range target {
					next[idx] = (target[idx] - pattern.value[idx]) >> 1
				}
				ans = min(ans, pattern.cost+2*helper(next))
			}
			memo[key] = ans
			return ans
		}
		return helper(target)
	}
	parse := func(s string) ([]int, [][]int, []int) {
		fs := strings.Fields(s)
		buttons := make([][]int, len(fs)-2)
		for idx, button := range fs[1 : len(fs)-1] {
			button = strings.Trim(button, "()")
			buttons[idx] = shared.IntSlice(strings.Split(button, ","))
		}
		ls := strings.Trim(fs[0], "[]")
		lights := make([]int, len(ls))
		for idx, c := range ls {
			if c == '#' {
				lights[idx] = 1
			}
		}
		targets := strings.Trim(fs[len(fs)-1], "{}")
		return lights, buttons, shared.IntSlice(strings.Split(targets, ","))
	}
	for _, line := range input {
		lights, buttons, target := parse(line)
		coeffs := makeCoeffs(buttons, len(target))
		patterns := makePatterns(coeffs)
		solution.Part1 += findParity(patterns, lights)
		solution.Part2 += findAddends(patterns, target)
	}
	return
}
