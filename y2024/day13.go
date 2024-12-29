package y2024

import (
	"adventofcode/shared"
	"regexp"
)

func Day13(input []string) (solution shared.Solution[int, int]) {
	re := regexp.MustCompile(`\d+`)
	increment := 10000000000000
	solver := func(A, B, P []int) int {
		det := A[0]*B[1] - A[1]*B[0]
		X := []int{
			(P[0]*B[1] - P[1]*B[0]) / det,
			(P[1]*A[0] - P[0]*A[1]) / det,
		}
		if A[0]*X[0]+B[0]*X[1] == P[0] && A[1]*X[0]+B[1]*X[1] == P[1] {
			return 3*X[0] + X[1]
		}
		return 0
	}
	for i := 0; i < len(input); i += 4 {
		A := shared.IntSlice(re.FindAllString(input[i], 2))
		B := shared.IntSlice(re.FindAllString(input[i+1], 2))
		P := shared.IntSlice(re.FindAllString(input[i+2], 2))
		solution.Part1 += solver(A, B, P)
		solution.Part2 += solver(A, B, []int{P[0] + increment, P[1] + increment})
	}
	return
}
