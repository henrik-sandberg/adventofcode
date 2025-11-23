package y2022

import (
	"adventofcode/shared"
	"strconv"
	"strings"
)

func Day10(input []string) (solution shared.Solution[int, string]) {
	solution.Part1 = day10_part1(input)
	solution.Part2 = day10_part2(input)
	return
}

func day10_part1(commands []string) (result int) {
	cycle := 1
	register := 1
	for _, cmd := range commands {
		cycle += 1
		if cycle%40 == 20 {
			result += cycle * register
		}
		if cmd == "noop" {
			continue
		}
		cmdArr := strings.Split(cmd, " ")
		val, _ := strconv.Atoi(cmdArr[1])
		register += val
		cycle += 1
		if cycle%40 == 20 {
			result += cycle * register
		}
	}
	return
}

func day10_part2(commands []string) string {
	width := 40
	cycle := 0
	sprite := 1
	var crt strings.Builder
	for _, cmd := range commands {
		crt.WriteString(getCrtChar(cycle, sprite, width))
		cycle += 1
		if cmd == "noop" {
			continue
		}
		crt.WriteString(getCrtChar(cycle, sprite, width))
		cycle += 1
		cmdArr := strings.Split(cmd, " ")
		val, _ := strconv.Atoi(cmdArr[1])
		sprite += val
	}
	res := "\n"
	for i, s := 0, crt.String(); i < len(s); i += width {
		res += s[i:i+width] + "\n"
	}
	return res
}

func getCrtChar(cycle int, sprite int, width int) string {
	if abs((cycle)%width-sprite) < 2 {
		return "#"
	} else {
		return "."
	}
}
