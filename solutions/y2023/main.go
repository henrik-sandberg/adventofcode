package y2023

import (
	"adventofcode/solutions/shared"
)

var Solutions = map[string]func([]string) shared.Solution[any, any]{
	"01": shared.WrapSolution(Day01),
	"02": shared.WrapSolution(Day02),
	"03": shared.WrapSolution(Day03),
	"04": shared.WrapSolution(Day04),
	"05": shared.WrapSolution(Day05),
	"06": shared.WrapSolution(Day06),
	"07": shared.WrapSolution(Day07),
	"08": shared.WrapSolution(Day08),
	"09": shared.WrapSolution(Day09),
	"10": shared.WrapSolution(Day10),
	"11": shared.WrapSolution(Day11),
	"12": shared.WrapSolution(Day12),
	"13": shared.WrapSolution(Day13),
	"14": shared.WrapSolution(Day14),
	"15": shared.WrapSolution(Day15),
	"16": shared.WrapSolution(Day16),
	"17": shared.WrapSolution(Day17),
	"18": shared.WrapSolution(Day18),
	"19": shared.WrapSolution(Day19),
	"20": shared.WrapSolution(Day20),
	"21": shared.WrapSolution(Day21),
	"23": shared.WrapSolution(Day23),
	"24": shared.WrapSolution(Day24),
	"25": shared.WrapSolution(Day25),
}

