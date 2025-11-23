package y2022

import (
	"adventofcode/solutions/shared"
	"fmt"
)

func Day03(input []string) (solution shared.Solution[int, int]) {
	solution.Part1 = day03_part1(input)
	solution.Part2 = day03_part2(input)
	return
}

func day03_part1(input []string) (result int) {
	for _, bag := range input {
		runeArray := []rune(bag)
		intersect := shared.Intersect(runeArray[:len(runeArray)/2], runeArray[len(runeArray)/2:])
		result += calculatePriority(intersect)
	}
	return
}

func day03_part2(input []string) (result int) {
	for i := 0; i < len(input); i += 3 {
		common := shared.Intersect([]rune(input[i]), []rune(input[i+1]))
		common = shared.Intersect(common, []rune(input[i+2]))
		result += calculatePriority(common)
	}
	return
}

func calculatePriority(chars []rune) int {
	if len(chars) != 1 {
		fmt.Printf("Invalid number of chars in %v\n", chars)
	}
	num := int(chars[0])
	if num >= 97 {
		num -= 96
	} else {
		num -= 38
	}
	return num
}
