package y2021

import (
	"adventofcode/solutions/shared"
)

var Solutions = map[string]func([]string) shared.Solution[any, any]{
	"01": shared.WrapSolution(Day01),
	"02": shared.WrapSolution(Day02),
	"03": shared.WrapSolution(Day03),
	"04": shared.WrapSolution(Day04),
}
