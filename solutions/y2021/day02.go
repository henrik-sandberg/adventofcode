package y2021

import (
	"strconv"
	"strings"

	"adventofcode/solutions/shared"
)

func Day02(input []string) (solution shared.Solution[int, int]) {
	type instruction struct {
		command string
		value   int
	}

	instructions := make([]instruction, len(input))
	for i, line := range input {
		parts := strings.Fields(line)
		val, _ := strconv.Atoi(parts[1])
		instructions[i] = instruction{command: parts[0], value: val}
	}

	var x1, y1 int
	var x2, y2, aim int
	for _, inst := range instructions {
		switch inst.command {
		case "forward":
			x1 += inst.value
			x2 += inst.value
			y2 += inst.value * aim
		case "up":
			y1 -= inst.value
			aim -= inst.value
		case "down":
			y1 += inst.value
			aim += inst.value
		}
	}
	solution.Part1 = x1 * y1
	solution.Part2 = x2 * y2
	return
}
