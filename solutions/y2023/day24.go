package y2023

import (
	"math/big"
	"strconv"
	"strings"

	"adventofcode/solutions/shared"
)

func Day24(input []string) (solution shared.Solution[int, int]) {
	return day24solve(input, 2e14, 4e14)
}

func day24solve(input []string, from, to float64) (solution shared.Solution[int, int]) {
	type equation struct {
		p, v []int
	}
	var equations []equation
	for _, line := range input {
		tmp := strings.Split(line, "@")
		equations = append(equations, equation{
			p: shared.IntSlice(strings.Split(tmp[0], ",")),
			v: shared.IntSlice(strings.Split(tmp[1], ",")),
		})
	}
	rat := func(x int) *big.Rat {
		return big.NewRat(int64(x), 1)
	}
	for eqs := range shared.Combinations(equations, 2) {
		T, err := shared.GaussianElimination([][]*big.Rat{
			{
				rat(eqs[0].v[0]),
				rat(-eqs[1].v[0]),
				rat(eqs[1].p[0] - eqs[0].p[0]),
			},
			{
				rat(eqs[0].v[1]),
				rat(-eqs[1].v[1]),
				rat(eqs[1].p[1] - eqs[0].p[1]),
			},
		})
		if err != nil {
			continue
		}
		t0, _ := T[0].Float64()
		t1, _ := T[1].Float64()
		x := float64(eqs[0].p[0]) + t0*float64(eqs[0].v[0])
		y := float64(eqs[0].p[1]) + t0*float64(eqs[0].v[1])
		if t0 >= 0 && t1 >= 0 &&
			from <= x && x <= to &&
			from <= y && y <= to {
			solution.Part1++
		}
	}
	generateCrossProductRows := func(a, b equation) [][]*big.Rat {
		var ret [][]*big.Rat
		dv := []int{a.v[0] - b.v[0], a.v[1] - b.v[1], a.v[2] - b.v[2]}
		dp := []int{a.p[0] - b.p[0], a.p[1] - b.p[1], a.p[2] - b.p[2]}
		for i := range 3 {
			// i = component along which we compute cross-product (X/Y/Z)
			// j, k = other two components in cyclic order
			j, k := (i+1)%3, (i+2)%3

			row := make([]*big.Rat, 7) // px, py, pz, vx, vy, vz, const

			// Coefficients for unknowns
			row[i] = rat(0)
			row[j] = rat(dv[k])
			row[k] = rat(-dv[j])
			row[i+3] = rat(0)
			row[j+3] = rat(dp[k])
			row[k+3] = rat(-dp[j])

			// constant = (a.p × a.v)[component] - (b.p × b.v)[component]
			row[6] = rat(a.p[j]*a.v[k] - a.p[k]*a.v[j] - (b.p[j]*b.v[k] - b.p[k]*b.v[j]))
			ret = append(ret, row)
		}
		return ret
	}

	mat := append(generateCrossProductRows(equations[0], equations[1]), generateCrossProductRows(equations[0], equations[2])...)
	solved, err := shared.GaussianElimination(mat)
	if err != nil {
		panic(err)
	}
	sum := new(big.Rat)
	sum.Add(solved[0], solved[1])
	sum.Add(sum, solved[2])
	solution.Part2, _ = strconv.Atoi(sum.FloatString(0))
	return
}
