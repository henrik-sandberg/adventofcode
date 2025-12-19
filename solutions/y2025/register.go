package y2025

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
	//	"11": shared.WrapSolution(Day11),
	//	"12": shared.WrapSolution(Day12),
}
