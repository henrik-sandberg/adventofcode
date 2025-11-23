package y2023

import (
	"adventofcode/solutions/shared"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Day01(input []string) (solution shared.Solution[int, int]) {
	digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	tr := func(s string) string {
		if i := slices.Index(digits, s); i != -1 {
			return strconv.Itoa(i)
		}
		return s
	}
	digitRegexp := regexp.MustCompile("\\d")
	digitAndDigitNameRegexp := regexp.MustCompile("\\d|" + strings.Join(digits[1:], "|"))
	for _, line := range input {
		if nums := digitRegexp.FindAllString(line, -1); len(nums) > 0 {
			n, _ := strconv.Atoi(nums[0] + nums[len(nums)-1])
			solution.Part1 += n
		}
		for _, digit := range digits {
			line = strings.ReplaceAll(line, digit, digit+digit[len(digit)-1:])
		}
		nums := digitAndDigitNameRegexp.FindAllString(line, -1)
		n, _ := strconv.Atoi(tr(nums[0]) + tr(nums[len(nums)-1]))
		solution.Part2 += n
	}
	return
}
