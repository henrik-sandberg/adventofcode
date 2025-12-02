package y2025

import (
	"strconv"
	"strings"

	"adventofcode/solutions/shared"
)

func Day02(input []string) (solution shared.Solution[int, int]) {
	isRepeating := func(s string) bool {
		length := len(s)
		for i := 1; i <= length/2; i++ {
			if length%i == 0 && strings.Repeat(s[:i], length/i) == s {
				return true
			}
		}
		return false
	}

	for rang := range strings.SplitSeq(input[0], ",") {
		nums := shared.IntSlice(strings.SplitN(rang, "-", 2))
		for n := nums[0]; n <= nums[1]; n++ {
			s := strconv.Itoa(n)
			half := len(s) / 2
			if s[:half] == s[half:] {
				solution.Part1 += n
			}
			if isRepeating(s) {
				solution.Part2 += n
			}
		}
	}
	return
}
