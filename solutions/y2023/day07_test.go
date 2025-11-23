package y2023

import (
	"strings"
	"testing"
)

func TestDay07(t *testing.T) {
	sampleInput := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	input := strings.Split(sampleInput, "\n")

	res := Day07(input)

	if expected := 6440; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 5905; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
