package y2023

import (
	"strings"
	"testing"
)

func TestDay25(t *testing.T) {
	sampleInput := strings.Split(strings.TrimSpace(`
jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr
	`), "\n")

	res := Day25(sampleInput)

	if expected := 54; expected != res.Part1 {
		t.Errorf("part1: expected %d, got %d", expected, res.Part1)
	}
}
