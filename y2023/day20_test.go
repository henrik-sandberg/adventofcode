package y2023

import (
	"strings"
	"testing"
)

func TestDay20(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`), "\n")

	res := Day20(sampleInput)

	if expected := 11687500; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 0; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
