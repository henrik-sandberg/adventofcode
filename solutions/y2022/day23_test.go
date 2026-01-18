package y2022

import (
	"strings"
	"testing"
)

func TestDay23(t *testing.T) {
	sampleInput := `....#..
..###.#
#...#.#
.#...##
#.###..
##.#.##
.#..#..`

	input := strings.Split(sampleInput, "\n")

	res := Day23(input)

	if expected := 110; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 20; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
