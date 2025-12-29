package y2022

import (
	"adventofcode/solutions/shared"
)

func Day03(input []string) (solution shared.Solution[int, int]) {
	priority := func(chars []rune) int {
		num := int(chars[0])
		if num >= 97 {
			return num - 96
		}
		return num - 38
	}
	for _, bag := range input {
		runeArray := []rune(bag)
		intersect := shared.Intersect(runeArray[:len(runeArray)/2], runeArray[len(runeArray)/2:])
		solution.Part1 += priority(intersect)
	}
	for i := 0; i < len(input); i += 3 {
		common := shared.Intersect([]rune(input[i]), []rune(input[i+1]))
		common = shared.Intersect(common, []rune(input[i+2]))
		solution.Part2 += priority(common)
	}
	return
}
