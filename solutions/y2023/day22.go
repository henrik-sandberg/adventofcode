package y2023

import (
	"cmp"
	"regexp"
	"slices"

	"adventofcode/solutions/shared"
)

func Day22(input []string) (solution shared.Solution[int, int]) {
	type point struct {
		x, y, z int
	}
	type brick struct {
		from point
		to   point
	}
	re := regexp.MustCompile(`\d+`)
	var stack []brick
	for _, line := range input {
		ints := shared.IntSlice(re.FindAllString(line, -1))
		stack = append(stack, brick{
			from: point{x: ints[0], y: ints[1], z: ints[2]},
			to:   point{x: ints[3], y: ints[4], z: ints[5]},
		})
	}
	slices.SortStableFunc(stack, func(a, b brick) int {
		return cmp.Compare(a.from.z, b.from.z)
	})
	drop := func(stack []brick, skip int) int {
		peaks := map[point]int{}
		falls := 0
		for ind, br := range stack {
			if ind == skip {
				continue
			}
			var area []point
			for x := br.from.x; x <= br.to.x; x++ {
				for y := br.from.y; y <= br.to.y; y++ {
					area = append(area, point{x: x, y: y})
				}
			}
			peak := 0
			for _, a := range area {
				peak = max(peaks[a]+1, peak)
			}
			h := br.to.z - br.from.z
			for _, a := range area {
				peaks[a] = peak + h
			}
			stack[ind].from.z = peak
			stack[ind].to.z = peak + h
			if peak < br.from.z {
				falls++
			}
		}
		return falls
	}
	drop(stack, -1)
	for i := range len(stack) {
		cop := slices.Clone(stack)
		res := drop(cop, i)
		if res == 0 {
			solution.Part1++
		}
		solution.Part2 += res
	}
	return
}
