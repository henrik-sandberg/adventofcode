package y2024

import (
	"adventofcode/shared"
	"slices"
	"strings"
)

func Day02(input []string) shared.Solution {
	part1 := 0
	part2 := 0
	for _, line := range input {
		nums := shared.IntSlice(strings.Fields(line))
		if isSafeLevel(nums) {
			part1++
			part2++
		} else {
			for ind := range nums {
				oneDropped := append([]int{}, nums[:ind]...)
				oneDropped = append(oneDropped, nums[ind+1:]...)
				if isSafeLevel(oneDropped) {
					part2++
					break
				}
			}
		}
	}
	return shared.Solution{Part1: part1, Part2: part2}
}

func isSafeLevel(nums []int) bool {
	diffs := calculateDiffs(nums)
	for _, diff := range diffs {
		if shared.Abs(diff) < 1 || shared.Abs(diff) > 3 {
			return false
		}
	}
	positives := make([]int, 0, len(diffs))
	for _, i := range diffs {
		if i >= 0 {
			positives = append(positives, i)
		}
	}
	negatives := make([]int, 0, len(diffs))
	for _, i := range diffs {
		if i <= 0 {
			negatives = append(negatives, i)
		}
	}
	return slices.Equal(positives, diffs) || slices.Equal(negatives, diffs)
}

func calculateDiffs(ints []int) []int {
	diffs := make([]int, len(ints)-1)
	for i := range ints[1:] {
		diffs[i] = ints[i+1] - ints[i]
	}
	return diffs
}
