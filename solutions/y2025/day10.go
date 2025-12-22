package y2025

import (
	"adventofcode/solutions/shared"
	"slices"
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
	part1 := func(buttons [][]int, target []int) int {
		for k := range len(buttons) + 1 {
			for comb := range shared.Combinations(buttons, k) {
				work := make([]int, len(target))
				for _, n := range slices.Concat(comb...) {
					work[n] = (work[n] + 1) & 1
				}
				if slices.Equal(target, work) {
					return k
				}
			}
		}
		return 0
	}
	part2 := func(buttons [][]int, target []int) int {
		var coeffs [][]int
		for _, button := range buttons {
			c := make([]int, len(target))
			for _, i := range button {
				c[i] = 1
			}
			coeffs = append(coeffs, c)
		}
		patterns := makePatterns(coeffs)
		memo := make(map[string]int)
		var helper func(target []int) int
		helper = func(target []int) int {
			key := encode(target)
			if v, ok := memo[key]; ok {
				return v
			}
			if !slices.ContainsFunc(target, func(i int) bool {
				return i != 0
			}) {
				return 0
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

	for _, line := range input {
		fs := strings.Fields(line)
		var buttons [][]int
		for _, s := range fs[1 : len(fs)-1] {
			buttons = append(buttons, shared.IntSlice(strings.Split(strings.Trim(s, "()"), ",")))
		}
		lights := strings.Trim(fs[0], "[]")
		target := make([]int, len(lights))
		for idx, c := range lights {
			if c == '#' {
				target[idx] = 1
			}
		}
		solution.Part1 += part1(buttons, target)
		solution.Part2 += part2(buttons, shared.IntSlice(strings.Split(strings.Trim(fs[len(fs)-1], "{}"), ",")))
	}
	return
}
