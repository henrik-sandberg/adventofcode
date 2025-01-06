package y2023

import (
	"adventofcode/shared"
	"slices"
	"strings"
)

func Day12(input []string) (solution shared.Solution[int, int]) {
	solver := func(chars []byte, groups []int) int {
		type State struct {
			GroupIndex     int
			GroupCount     int
			RequireWorking bool
		}
		states := map[State]int{{}: 1}
		for _, char := range chars {
			newstates := map[State]int{}
			for state, val := range states {
				if (char == '#' || char == '?') && state.GroupIndex < len(groups) && !state.RequireWorking {
					if char == '?' && state.GroupCount == 0 {
						newstates[state] += val
					}
					state.GroupCount++
					if state.GroupCount == groups[state.GroupIndex] {
						state = State{state.GroupIndex + 1, 0, true}
					}
					newstates[state] += val
				} else if (char == '.' || char == '?') && state.GroupCount == 0 {
					state.RequireWorking = false
					newstates[state] += val
				}
			}
			states = newstates
		}
		res := 0
		for state, v := range states {
			if state.GroupIndex == len(groups) {
				res += v
			}
		}
		return res
	}
	for _, inpt := range input {
		tmp := strings.Fields(inpt)
		groups := shared.IntSlice(strings.Split(tmp[1], ","))
		solution.Part1 += solver(
			[]byte(tmp[0]),
			groups,
		)
		records := strings.Join(slices.Repeat(tmp[:1], 5), "?")
		solution.Part2 += solver(
			[]byte(records),
			slices.Repeat(groups, 5),
		)
	}
	return
}
