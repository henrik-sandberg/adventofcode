package shared

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

// Solution holds the results for Part 1 and Part 2 of a puzzle.
// It uses generics to allow for any result type.
type Solution[T1, T2 any] struct {
	Part1 T1
	Part2 T2
}

type Solver func([]string) Solution[any, any]

// WrapSolution wraps a specific solution function (e.g., one returning Solution[int, int])
// into a generic one that returns Solution[any, any].
func WrapSolution[T1, T2 any](f func([]string) Solution[T1, T2]) Solver {
	return func(input []string) Solution[any, any] {
		sol := f(input)
		return Solution[any, any]{Part1: sol.Part1, Part2: sol.Part2}
	}
}

func Run[T1, T2 any](f func([]string) Solution[T1, T2], inputReader io.Reader, outputWriter io.Writer) error {
	input, err := readLines(inputReader)
	if err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	start := time.Now()
	solution := f(input)
	elapsed := time.Since(start)

	fmt.Fprintf(outputWriter, "Part1: %v\n", solution.Part1)
	fmt.Fprintf(outputWriter, "Part2: %v\n", solution.Part2)
	fmt.Fprintf(outputWriter, "Solver completed in %d ms\n", elapsed.Milliseconds())
	return nil
}

func readLines(inputReader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading from input: %w", err)
	}
	return lines, nil
}
