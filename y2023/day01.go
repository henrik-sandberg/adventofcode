package y2023

import (
	"adventofcode/shared"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day01(input []string) (solution shared.Solution[int, int]) {
	re := regexp.MustCompile("\\d")
	for _, line := range input {
		all := re.FindAllString(line, -1)
		if len(all) > 0 {
			n, _ := strconv.Atoi(all[0] + all[len(all)-1])
			solution.Part1 += n
		}
	}
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	re = regexp.MustCompile("\\d|" + strings.Join(digits[1:], "|"))
	tr := func(s string) string {
		if i := slices.Index(digits, s); i != -1 {
			return strconv.Itoa(i)
		}
		return s
	}
	for _, line := range input {
		for _, digit := range digits {
			line = strings.ReplaceAll(line, digit, digit+digit[len(digit)-1:])
		}
		all := re.FindAllString(line, -1)
		n, _ := strconv.Atoi(tr(all[0]) + tr(all[len(all)-1]))
		solution.Part2 += n
	}
	return
}
