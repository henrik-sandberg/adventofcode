package y2022

import (
	"strings"

	"adventofcode/solutions/shared"
)

func Day02(input []string) (solution shared.Solution[int, int]) {
	type gameOption struct {
		name  string
		beats string
		loses string
		value int
	}
	gameOptions := map[string]gameOption{
		"rock":    {name: "rock", beats: "scissor", loses: "paper", value: 1},
		"paper":   {name: "paper", beats: "rock", loses: "scissor", value: 2},
		"scissor": {name: "scissor", beats: "paper", loses: "rock", value: 3},
	}
	score := func(us, they gameOption) int {
		// Score: 0 if lost, 3 if draw, 6 if win
		// plus value
		if us.name == they.name {
			return 3 + us.value
		}
		if they.name == us.beats {
			return 6 + us.value
		}
		return us.value
	}
	part1 := func() int {
		result := 0
		mappings := map[string]string{
			"A": "rock",
			"B": "paper",
			"C": "scissor",
			"X": "rock",
			"Y": "paper",
			"Z": "scissor",
		}
		for _, round := range input {
			a, b, _ := strings.Cut(round, " ")
			result += score(gameOptions[mappings[a]], gameOptions[mappings[b]])
		}
		return result
	}
	part2 := func() int {
		result := 0
		mappings := map[string]string{
			"A": "rock",
			"B": "paper",
			"C": "scissor",
		}

		for _, round := range input {
			a, b, _ := strings.Cut(round, " ")
			they := gameOptions[mappings[a]]
			var us gameOption
			switch b {
			case "X":
				us = gameOptions[they.beats]
			case "Y":
				us = gameOptions[they.name]
			default:
				us = gameOptions[they.loses]
			}
			result += score(us, they)
		}
		return result
	}

	solution.Part1 = part1()
	solution.Part2 = part2()
	return
}
