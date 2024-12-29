package shared

import "math"

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
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

func lcm(n ...int) int {
	if len(n) < 2 {
		panic("Can only calculate LCM of at least 2 numbers")
	}
	res := n[0] * (n[1] / gcd(n[0], n[1]))
	for _, v := range n[2:] {
		res = lcm(res, v)
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

// Returns 1 if true, else 0
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
