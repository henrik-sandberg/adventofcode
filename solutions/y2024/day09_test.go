package y2024

import (
	"testing"
)

func TestDay09(t *testing.T) {
	input := []string{"2333133121414131402"}

	res := Day09(input)

	if expected := 1928; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 2858; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
