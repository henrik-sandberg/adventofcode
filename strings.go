package main

import "strconv"

func Reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func AllIndex(s string, r rune) []int {
	ret := make([]int, 0, len(s))
	for i, v := range []rune(s) {
		if v == r {
			ret = append(ret, i)
		}
	}
	return ret
}

func ToInts(input []string) []int {
	ints := make([]int, len(input), len(input))
	for i, s := range input {
		value, _ := strconv.Atoi(s)
		ints[i] = value
	}
	return ints
}
