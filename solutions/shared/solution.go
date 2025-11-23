package shared

// Solution holds the results for Part 1 and Part 2 of a puzzle.
// It uses generics to allow for any result type.
type Solution[T1, T2 any] struct {
	Part1 T1
	Part2 T2
}

// WrapSolution wraps a specific solution function (e.g., one returning Solution[int, int])
// into a generic one that returns Solution[any, any].
func WrapSolution[T1, T2 any](f func([]string) Solution[T1, T2]) func([]string) Solution[any, any] {
	return func(input []string) Solution[any, any] {
		sol := f(input)
		return Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	}
}
