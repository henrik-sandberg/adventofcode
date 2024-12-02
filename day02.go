package main

import (
	"fmt"
	"strings"
)

func Day02(input []string) {
	part1 := 0
	part2 := 0
	for _, line := range input {
		nums := ToInts(strings.Fields(line))
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
	fmt.Println(part1)
	fmt.Println(part2)
}

func isSafeLevel(nums []int) bool {
	diffs := calculateDiffs(nums)
	for _, diff := range diffs {
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}
	}
	positives := len(Filter(diffs, func(i int) bool {
		return i >= 0
	}))
	negatives := len(Filter(diffs, func(i int) bool {
		return i <= 0
	}))
	return positives == len(diffs) || negatives == len(diffs)

}

func calculateDiffs(ints []int) []int {
	diffs := make([]int, len(ints)-1, len(ints)-1)
	for i := range ints[1:] {
		diffs[i] = ints[i+1] - ints[i]
	}
	return diffs
}
