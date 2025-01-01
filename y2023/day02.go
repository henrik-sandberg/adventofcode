package y2023

import (
	"adventofcode/shared"
	"strconv"
	"strings"
)

func Day02(input []string) (solution shared.Solution[int, int]) {
	for i, line := range input {
		samples := parseGame(line)
		if canUseLimitedNumberOfCubes(samples) {
			solution.Part1 += i + 1
		}
		solution.Part2 += minPossibleCubesProduct(samples)
	}
	return
}

func canUseLimitedNumberOfCubes(samples []map[string]int) bool {
	for _, sample := range samples {
		if sample["red"] > 12 || sample["green"] > 13 || sample["blue"] > 14 {
			return false
		}
	}
	return true
}

func minPossibleCubesProduct(samples []map[string]int) int {
	maxUsed := map[string]int{}
	for _, sample := range samples {
		for k, v := range sample {
			maxUsed[k] = max(maxUsed[k], v)
		}
	}
	res := 1
	for _, v := range maxUsed {
		res *= v
	}
	return res
}

func parseGame(raw string) []map[string]int {
	samples := []map[string]int{}
	for _, sample := range strings.Split(strings.SplitN(raw, ": ", 2)[1], "; ") {
		drawnCubes := map[string]int{}
		for _, cubes := range strings.Split(sample, ", ") {
			cube := strings.Split(cubes, " ")
			drawnCubes[cube[1]], _ = strconv.Atoi(cube[0])
		}
		samples = append(samples, drawnCubes)
	}
	return samples
}
