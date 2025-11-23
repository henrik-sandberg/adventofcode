package y2023

import (
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {

	sampleInput := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

	input := strings.Split(sampleInput, "\n")

	res := Day03(input)

	if expected := 4361; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 467835; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
