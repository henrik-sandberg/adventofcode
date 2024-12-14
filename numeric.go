package main

import "math"

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func sum(n ...int) (res int) {
	for _, v := range n {
		res += v
	}
	return res
}

func multiply(n ...int) int {
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

func stddev(arr []int) float64 {
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

func positiveMod(i, mod int) int {
	return (i%mod + mod) % mod
}
