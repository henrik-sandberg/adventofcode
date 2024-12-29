package y2024

import (
	"strings"
	"testing"
)

func TestDay20(t *testing.T) {
	sampleInput := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

	input := strings.Split(sampleInput, "\n")

	res := Day20Solver(input, 2)

	if expected := 44; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	res = Day20Solver(input, 50)
	if expected := 285; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
