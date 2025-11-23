package y2023

import (
	"adventofcode/shared"
	"strings"
)

func Day24(input []string) (solution shared.Solution[int, int]) {
	day24 := day24{input}
	solution.Part1 = day24.part1(2e14, 4e14)
	return
}

type day24 struct {
	input []string
}

func (d *day24) part1(from, to float64) (res int) {
	type equation struct {
		p, v []int
	}
	equations := []equation{}
	for _, line := range d.input {
		tmp := strings.Split(line, "@")
		equations = append(equations, equation{
			shared.IntSlice(strings.Split(tmp[0], ",")),
			shared.IntSlice(strings.Split(tmp[1], ",")),
		})
	}
	for eqs := range shared.Combinations(equations, 2) {
		A := [][]int{
			{eqs[0].v[0], -eqs[1].v[0]},
			{eqs[0].v[1], -eqs[1].v[1]},
		}
		det := A[0][0]*A[1][1] - A[0][1]*A[1][0]
		if det == 0 {
			continue
		}
		B := []int{
			eqs[1].p[0] - eqs[0].p[0],
			eqs[1].p[1] - eqs[0].p[1],
		}
		T := []float64{
			float64(B[0]*A[1][1]-B[1]*A[0][1]) / float64(det),
			float64(B[1]*A[0][0]-B[0]*A[1][0]) / float64(det),
		}
		X := []float64{
			float64(eqs[0].p[0]) + T[0]*float64(eqs[0].v[0]),
			float64(eqs[1].p[0]) + T[1]*float64(eqs[1].v[0]),
		}
		Y := []float64{
			float64(eqs[0].p[1]) + T[0]*float64(eqs[0].v[1]),
			float64(eqs[1].p[1]) + T[1]*float64(eqs[1].v[1]),
		}
		if T[0] >= 0 && T[1] >= 0 &&
			from <= X[0] && X[0] <= to &&
			from <= Y[0] && Y[0] <= to &&
			from <= X[1] && X[1] <= to &&
			from <= Y[1] && Y[1] <= to {
			res++
		}
	}
	return res
}
