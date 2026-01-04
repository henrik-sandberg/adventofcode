package y2022

import (
	"strings"
	"testing"
)

func TestDay17(t *testing.T) {
	sampleInput := `>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`

	input := strings.Split(sampleInput, "\n")

	res := Day17(input)

	if expected := 3068; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 1514285714288; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
