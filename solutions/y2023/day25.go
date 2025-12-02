package y2023

import (
	"cmp"
	"errors"
	"maps"
	"slices"
	"strings"

	"adventofcode/solutions/shared"
)

func Day25(input []string) (solution shared.Solution[int, any]) {
	G := map[string][]string{}
	for _, line := range input {
		fs := strings.Split(line, ": ")
		u := fs[0]
		for _, v := range strings.Fields(fs[1]) {
			G[u] = append(G[u], v)
			G[v] = append(G[v], u)
		}
	}
	solver := func() (int, error) {
		S := map[string]bool{}
		for k := range G {
			S[k] = true
		}
		countDisjoint := func(set []string) int {
			n := len(set)
			for _, e := range set {
				if S[e] {
					n--
				}
			}
			return n
		}
		sum := func() int {
			res := 0
			for k := range S {
				res += countDisjoint(G[k])
			}
			return res
		}
		for sum() != 3 {
			s := slices.MaxFunc(slices.Collect(maps.Keys(S)), func(a, b string) int {
				return cmp.Compare(countDisjoint(G[a]), countDisjoint(G[b]))
			})
			delete(S, s)
			if len(S) == 0 {
				return 0, errors.New("failed to resolve")
			}
		}
		return len(S) * countDisjoint(slices.Collect(maps.Keys(G))), nil
	}

	for {
		var err error
		solution.Part1, err = solver()
		if err == nil {
			break
		}
	}
	return
}
