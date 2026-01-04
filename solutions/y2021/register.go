package y2021

import (
	"adventofcode/solutions/shared"
)

var Solvers = map[string]shared.Solver{
	"01": shared.WrapSolution(Day01),
	"02": shared.WrapSolution(Day02),
	"03": shared.WrapSolution(Day03),
	"04": shared.WrapSolution(Day04),
	"05": shared.WrapSolution(Day05),
}
