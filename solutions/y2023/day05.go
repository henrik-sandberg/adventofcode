package y2023

import (
	"adventofcode/solutions/shared"
	"math"
	"strings"
)

func Day05(input []string) (solution shared.Solution[int, int]) {
	type Mapping struct {
		Dst    int
		Src    int
		Length int
	}
	type Seed struct {
		Value  int
		Length int
	}
	mappings := [][]Mapping{}
	for _, line := range input[2:] {
		if line == "" {
			continue
		}
		if strings.HasSuffix(line, "map:") {
			mappings = append(mappings, []Mapping{})
			continue
		}
		imap := shared.IntSlice(strings.Fields(line))
		m := &mappings[len(mappings)-1]
		*m = append(*m, Mapping{imap[0], imap[1], imap[2]})
	}

	solver := func(seeds []Seed) int {
		for _, mapping := range mappings {
			var next []Seed
			for _, seed := range seeds {
			loop:
				for seed.Length > 0 {
					for _, m := range mapping {
						delta := seed.Value - m.Src
						if 0 <= delta && delta < m.Length {
							length := min(m.Length-delta, seed.Length)
							next = append(next, Seed{m.Dst + delta, length})
							seed.Value += length
							seed.Length -= length
							continue loop
						}
					}
					// Did not find any mapping. Default is map to same
					next = append(next, seed)
					break
				}
			}
			seeds = next
		}
		best := math.MaxInt
		for _, seed := range seeds {
			best = min(best, seed.Value)
		}
		return best
	}
	inp := shared.IntSlice(strings.Fields(input[0])[1:])
	var seeds []Seed
	for _, val := range inp {
		seeds = append(seeds, Seed{val, 1})
	}
	solution.Part1 = solver(seeds)

	seeds = nil
	for i := 0; i < len(inp); i += 2 {
		seeds = append(seeds, Seed{inp[i], inp[i+1]})
	}
	solution.Part2 = solver(seeds)
	return
}
