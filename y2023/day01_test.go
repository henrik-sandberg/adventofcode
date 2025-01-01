package y2023

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	input := strings.Split(sampleInput, "\n")

	res := Day01(input)

	if expected := 142; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	sampleInput = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

	input = strings.Split(sampleInput, "\n")

	res = Day01(input)

	if expected := 281; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
