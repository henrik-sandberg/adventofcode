package y2022

import (
	"adventofcode/solutions/shared"
	"strconv"
	"strings"
)

func Day04(input []string) (solution shared.Solution[int, int]) {
	for _, line := range input {
		first, second := toIntPairs(line)
		solution.Part1 += shared.BoolToInt(first[0] >= second[0] && first[1] <= second[1] || second[0] >= first[0] && second[1] <= first[1])
		solution.Part2 += shared.BoolToInt(first[0] >= second[0] && first[0] <= second[1] || first[1] >= second[0] && first[1] <= second[1] || second[0] >= first[0] && second[0] <= first[1] || second[1] >= first[0] && second[1] <= first[1])
	}
	return
}

func toIntPairs(s string) ([]int, []int) {
	pair := strings.Split(s, ",")

	first := strings.Split(pair[0], "-")
	first_low, _ := strconv.Atoi(first[0])
	first_high, _ := strconv.Atoi(first[1])

	second := strings.Split(pair[1], "-")
	second_low, _ := strconv.Atoi(second[0])
	second_high, _ := strconv.Atoi(second[1])

	return []int{first_low, first_high}, []int{second_low, second_high}
}
