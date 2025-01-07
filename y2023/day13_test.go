package y2023

import (
	"strings"
	"testing"
)

func TestDay13(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
	`), "\n")

	res := Day13(sampleInput)

	if expected := 405; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 400; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
