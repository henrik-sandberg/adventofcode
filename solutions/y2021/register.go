package y2021

import (
	"adventofcode/solutions/shared"
)

var Solutions = map[string]func([]string) shared.Solution[any, any]{
	"01": shared.WrapSolution(Day01),
}
