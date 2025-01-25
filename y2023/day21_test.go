package y2023

import (
	"strings"
	"testing"
)

func TestDay21(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
	`), "\n")

	day21 := day21{sampleInput}

	res := day21.part1(6)
	if expected := 16; expected != res {
		t.Errorf("part1: expected %d, got %d", expected, res)
	}
}
