package y2023

import (
	"strings"
	"testing"
)

func TestDay14(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
	`), "\n")

	res := Day14(sampleInput)

	if expected := 136; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 64; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
