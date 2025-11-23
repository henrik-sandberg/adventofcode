package y2023

import (
	"strings"
	"testing"
)

func TestDay17(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533
	`), "\n")

	res := Day17(sampleInput)

	if expected := 102; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	if expected := 94; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
