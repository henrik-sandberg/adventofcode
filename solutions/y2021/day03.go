package y2021

import (
	"slices"
	"strconv"

	"adventofcode/solutions/shared"
)

func Day03(input []string) (solution shared.Solution[int, int]) {
	if len(input) == 0 {
		return
	}
	width := len(input[0])
	count := func(slice []string, ind int, search byte) int {
		ret := 0
		for _, line := range slice {
			if line[ind] == search {
				ret++
			}
		}
		return ret
	}
	ones := make([]int, width)
	for i := range width {
		ones[i] = count(input, i, '1')
	}
	gamma := 0
	for i, n := range ones {
		if n > len(input)/2 {
			gamma |= 1 << (width - i - 1)
		}
	}
	epsilon := gamma ^ (1<<len(ones) - 1)
	solution.Part1 = gamma * epsilon

	tmp := slices.Clone(input)
	for len(tmp) > 1 {
		for i := range width {
			filter := byte('1')
			if ones := count(tmp, i, '1'); ones < len(tmp)-ones {
				filter = '0'
			}
			tmp = slices.DeleteFunc(tmp, func(s string) bool {
				return s[i] != filter
			})
		}
	}
	res, _ := strconv.ParseInt(tmp[0], 2, 0)
	solution.Part2 = int(res)

	tmp = slices.Clone(input)
	for len(tmp) > 1 {
		for i := range width {
			filter := byte('0')
			if zeroes := count(tmp, i, '0'); zeroes > len(tmp)-zeroes {
				filter = '1'
			}
			tmp = slices.DeleteFunc(tmp, func(s string) bool {
				return s[i] != filter
			})
			if len(tmp) == 1 {
				break
			}
		}
	}
	res, _ = strconv.ParseInt(tmp[0], 2, 0)
	solution.Part2 *= int(res)
	return
}
