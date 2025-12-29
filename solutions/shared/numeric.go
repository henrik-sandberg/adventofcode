package shared

import (
	"fmt"
	"math"
	"math/big"
)

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func Sign(i int) int {
	switch {
	case i < 0:
		return -1
	case i > 0:
		return 1
	default:
		return 0
	}
}

func Sum(n ...int) (res int) {
	for _, v := range n {
		res += v
	}
	return res
}

func Product(n ...int) int {
	product := n[0]
	for _, v := range n[1:] {
		product *= v
	}
	return product
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate the least common multiple
func LCM(n ...int) int {
	if len(n) < 2 {
		panic("Can only calculate LCM of at least 2 numbers")
	}
	res := n[0] * (n[1] / gcd(n[0], n[1]))
	for _, v := range n[2:] {
		res = LCM(res, v)
	}
	return res
}

func Stddev(arr []int) float64 {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	variance := 0.0
	for _, v := range arr {
		diff := float64(v) - mean
		variance += diff * diff
	}
	variance /= float64(len(arr))
	return math.Sqrt(variance)
}

func PositiveMod(i, mod int) int {
	return (i%mod + mod) % mod
}

// Solves an n×(n+1) augmented matrix in-place.
// Returns solution vector of length n.
func GaussianElimination(mat [][]*big.Rat) ([]*big.Rat, error) {
	n := len(mat)

	zero := big.NewRat(0, 1)

	for col := 0; col < n; col++ {
		// Find pivot
		pivot := col
		for pivot < n && mat[pivot][col].Cmp(zero) == 0 {
			pivot++
		}
		if pivot == n {
			return nil, fmt.Errorf("singular matrix")
		}

		// Swap rows
		mat[col], mat[pivot] = mat[pivot], mat[col]

		// Normalize pivot row
		pivotVal := new(big.Rat).Set(mat[col][col])
		for j := col; j <= n; j++ {
			mat[col][j].Quo(mat[col][j], pivotVal)
		}

		// Eliminate other rows
		for i := 0; i < n; i++ {
			if i == col {
				continue
			}
			factor := new(big.Rat).Set(mat[i][col])
			if factor.Cmp(zero) == 0 {
				continue
			}
			for j := col; j <= n; j++ {
				tmp := new(big.Rat).Mul(factor, mat[col][j])
				mat[i][j].Sub(mat[i][j], tmp)
			}
		}
	}

	solution := make([]*big.Rat, n)
	for i := range n {
		solution[i] = mat[i][n]
	}
	return solution, nil
}
