package main

import (
	"fmt"
	"regexp"
	"slices"
)

func Day13(input []string) {
	re := regexp.MustCompile(`\d+`)
	p2_increment := 10000000000000
	part1 := 0
	part2 := 0
	for i := 0; i < len(input); i += 4 {
		A := ToInts(re.FindAllString(input[i], 2))
		B := ToInts(re.FindAllString(input[i+1], 2))
		P := ToInts(re.FindAllString(input[i+2], 2))
		M := [][]int{
			{A[0], B[0]},
			{A[1], B[1]},
		}
		p1 := clawMachine(M, P)
		p2 := clawMachine(M, []int{P[0] + p2_increment, P[1] + p2_increment})
		part1 += 3*p1[0] + p1[1]
		part2 += 3*p2[0] + p2[1]
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func clawMachine(M [][]int, P []int) []int {
	X := []int{
		determinant([][]int{
			{P[0], M[0][1]},
			{P[1], M[1][1]},
		}) / determinant(M),
		determinant([][]int{
			{M[0][0], P[0]},
			{M[1][0], P[1]},
		}) / determinant(M),
	}
	calculated := []int{
		M[0][0]*X[0] + M[0][1]*X[1],
		M[1][0]*X[0] + M[1][1]*X[1],
	}
	if slices.Equal(P, calculated) {
		return X
	}
	return []int{0, 0}
}

func determinant(arr [][]int) int {
	return arr[0][0]*arr[1][1] - arr[0][1]*arr[1][0]
}
