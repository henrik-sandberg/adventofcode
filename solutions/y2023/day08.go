package y2023

import (
	"adventofcode/solutions/shared"
	"strings"
)

type tuple struct {
	left, right string
}

func Day08(input []string) (solution shared.Solution[int, int]) {
	dirs := input[0]
	network := parseNetwork(input[2:])
	solution.Part1 = stepsUntilTarget("AAA", "ZZZ", dirs, network)
	var paths []int
	for k := range network {
		if strings.HasSuffix(k, "A") {
			paths = append(paths, stepsUntilTarget(k, "Z", dirs, network))
		}
	}
	solution.Part2 = shared.LCM(paths...)
	return
}

func stepsUntilTarget(start, targetSuffix, dirs string, network map[string]tuple) (res int) {
	next := start
	for !strings.HasSuffix(next, targetSuffix) {
		dir := string(dirs[res%len(dirs)])
		if dir == "L" {
			next = network[next].left
		} else {
			next = network[next].right
		}
		res++
	}
	return
}

func parseNetwork(raw []string) map[string]tuple {
	network := make(map[string]tuple, len(raw))
	for _, line := range raw {
		split := strings.Split(line, " = ")
		vals := split[1][1 : len(split[1])-1]
		lr := strings.Split(vals, ", ")
		network[split[0]] = tuple{lr[0], lr[1]}
	}
	return network
}
