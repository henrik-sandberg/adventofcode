package shared

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

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
	lines := []string{}
	scanner := bufio.NewScanner(inputReader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading from input: %w", err)
	}
	return lines, nil
}
