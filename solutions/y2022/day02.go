package y2022

import (
	"strings"

	"adventofcode/solutions/shared"
)

type gameOption struct {
	name  string
	beats string
	loses string
	value int
}

func Day02(input []string) (solution shared.Solution[int, int]) {
	rock := gameOption{name: "rock", beats: "scissor", loses: "paper", value: 1}
	paper := gameOption{name: "paper", beats: "rock", loses: "scissor", value: 2}
	scissor := gameOption{name: "scissor", beats: "paper", loses: "rock", value: 3}

	gameOptions := make(map[string]gameOption)
	gameOptions[rock.name] = rock
	gameOptions[paper.name] = paper
	gameOptions[scissor.name] = scissor

	solution.Part1 = day02_part1(input, gameOptions)
	solution.Part2 = day02_part2(input, gameOptions)
	return
}

func day02_part1(input []string, gameOptions map[string]gameOption) (score int) {
	mappings := make(map[string]string)
	mappings["A"] = "rock"
	mappings["B"] = "paper"
	mappings["C"] = "scissor"

	mappings["X"] = "rock"
	mappings["Y"] = "paper"
	mappings["Z"] = "scissor"

	for _, round := range input {
		s := strings.Split(round, " ")
		score += calculate_score(gameOptions[mappings[s[1]]], gameOptions[mappings[s[0]]])
	}
	return
}

func day02_part2(input []string, gameOptions map[string]gameOption) (score int) {
	mappings := make(map[string]string)
	mappings["A"] = "rock"
	mappings["B"] = "paper"
	mappings["C"] = "scissor"

	for _, round := range input {
		s := strings.Split(round, " ")
		they := gameOptions[mappings[s[0]]]
		var us gameOption
		switch s[1] {
		case "X":
			us = gameOptions[they.beats]
		case "Y":
			us = gameOptions[they.name]
		default:
			us = gameOptions[they.loses]
		}
		score += calculate_score(us, they)
	}
	return
}

func calculate_score(us gameOption, they gameOption) int {
	// Score: 0 if lost, 3 if draw, 6 if win
	// plus value
	if us.name == they.name {
		return 3 + us.value
	} else if they.name == us.beats {
		return 6 + us.value
	} else {
		return us.value
	}
}
