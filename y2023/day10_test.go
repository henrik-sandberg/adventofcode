package y2023

import (
	"strings"
	"testing"
)

func TestDay10(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
..F7.
.FJ|.
SJ.L7
|F--J
LJ...
`), "\n")

	res := Day10(sampleInput)

	if expected := 8; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	sampleInput = strings.Split(strings.TrimSpace(`
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
`), "\n")

	res = Day10(sampleInput)

	if expected := 10; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
