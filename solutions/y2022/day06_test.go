package y2022

import (
	"strings"
	"testing"
)

func TestDay06(t *testing.T) {
	sampleInput := `mjqjpqmgbljsphdztnvjfqwrcgsmlb`

	input := strings.Split(sampleInput, "\n")

	res := Day06(input)

	if expected := 7; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 19; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
