package y2024

import (
	"strings"
	"testing"
)

func TestDay22(t *testing.T) {
	sampleInput := `1
10
100
2024`

	input := strings.Split(sampleInput, "\n")

	res := Day22(input)

	if expected := 37327623; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	sampleInput = `1
2
3
2024`

	input = strings.Split(sampleInput, "\n")
	res = Day22(input)

	if expected := 23; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
