package y2022

import (
	"strconv"
	"strings"

	"adventofcode/solutions/shared"
)

func Day05(input []string) (solution shared.Solution[string, string]) {
	solution.Part1 = day05_part1(input)
	solution.Part2 = day05_part2(input)
	return
}

func createStacks(input []string) (int, map[int]string) {
	stacks := make(map[int]string)
	for lineNumber, line := range input {
		if strings.Contains(line, "[") {
			for i := 0; i*4 < len(line); i++ {
				if crate := line[i*4 : i*4+3]; strings.TrimSpace(crate) != "" {
					stacks[i+1] = string(crate[1]) + stacks[i+1]
				}
			}
		}
		if line == "" {
			return lineNumber, stacks
		}
	}
	return 0, stacks
}

func day05_part1(input []string) string {
	i, stacks := createStacks(input)
	for _, move := range input[i+1:] {
		c, f, t := parseCommand(move)
		source := stacks[f]
		stacks[t] = stacks[t] + shared.Reverse(source[len(source)-c:])
		stacks[f] = source[:len(source)-c]
	}
	return getResult(stacks)
}

func day05_part2(input []string) string {
	i, stacks := createStacks(input)
	for _, move := range input[i+1:] {
		c, f, t := parseCommand(move)
		source := stacks[f]
		stacks[t] = stacks[t] + source[len(source)-c:]
		stacks[f] = source[:len(source)-c]
	}
	return getResult(stacks)
}

func parseCommand(command string) (int, int, int) {
	cmd := strings.Split(command, " ")
	count, _ := strconv.Atoi(cmd[1])
	from, _ := strconv.Atoi(cmd[3])
	to, _ := strconv.Atoi(cmd[5])
	return count, from, to
}

func getResult(stacks map[int]string) string {
	result := ""
	for i := range len(stacks) {
		s := stacks[i+1]
		result += s[len(s)-1:]
	}
	return result
}
