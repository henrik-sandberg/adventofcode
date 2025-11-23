package y2024

import (
	"testing"
)

func TestDay03(t *testing.T) {
	res := Day03([]string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"})
	if expected := 161; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	res = Day03([]string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"})
	if expected := 161; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}
	if expected := 48; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
