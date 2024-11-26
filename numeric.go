package main

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(n ...int) int {
	min := n[0]
	for _, v := range n[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func max(n ...int) int {
	max := n[0]
	for _, v := range n[1:] {
		if v > max {
			max = v
		}
	}
	return max
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
              a, b = b, a % b
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
