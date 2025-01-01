package y2022

import (
	"strings"
	"testing"
)

func TestDay01(t *testing.T) {
	sampleInput := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	input := strings.Split(sampleInput, "\n")

	res := Day01(input)

	if expected := 24000; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 45000; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
