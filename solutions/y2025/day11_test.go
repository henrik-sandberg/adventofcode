package y2025

import (
	"strings"
	"testing"
)

func TestDay11(t *testing.T) {
	sampleInput := `aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out`

	input := strings.Split(sampleInput, "\n")

	res := Day11(input)

	if expected := 5; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}

	sampleInput = `svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out`

	input = strings.Split(sampleInput, "\n")

	res = Day11(input)

	if expected := 2; expected != res.Part2 {
		t.Errorf("part2: expected %d, got %d", expected, res.Part2)
	}
}
