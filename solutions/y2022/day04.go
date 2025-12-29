package y2022

import (
	"adventofcode/solutions/shared"
	"strconv"
	"strings"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	parse := func(s string) (int, int) {
		first, second, _ := strings.Cut(s, "-")
		low, _ := strconv.Atoi(first)
		high, _ := strconv.Atoi(second)
		return low, high
	}
	for _, line := range input {
		first, second, _ := strings.Cut(line, ",")
		fl, fh := parse(first)
		sl, sh := parse(second)
		if fl >= sl && fh <= sh || sl >= fl && sh <= fh {
			solution.Part1++
		}
		if fl >= sl && fl <= sh || fh >= sl && fh <= sh || sl >= fl && sl <= fh || sh >= fl && sh <= fh {
			solution.Part2++
		}
	}
	return
}
