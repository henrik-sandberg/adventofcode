package y2022

import (
	"strings"
	"testing"
)

func TestDay22(t *testing.T) {
	sampleInput := `        ...#
        .#..
        #...
        ....
...#.......#
........#...
..#....#....
..........#.
        ...#....
        .....#..
        .#......
        ......#.

10R5L5R10L4R5L5`

	input := strings.Split(sampleInput, "\n")

	res := Day22(input)

	if expected := 6032; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 0; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
