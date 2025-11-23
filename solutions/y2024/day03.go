package y2024

import (
	"adventofcode/solutions/shared"
	"regexp"
	"strconv"
	"strings"
)

func Day03(input []string) (solution shared.Solution[int, int]) {
	inp := strings.Join(input, "")
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)|(do|don't)\(\)`)
	enabled := true
	for _, match := range re.FindAllStringSubmatch(inp, -1) {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		solution.Part1 += a * b
		if match[3] == "do" {
			enabled = true
		} else if match[3] == "don't" {
			enabled = false
		}
		if enabled {
			solution.Part2 += a * b
		}
	}
	return
}
