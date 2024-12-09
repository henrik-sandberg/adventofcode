package main

import (
	"fmt"
)

func Day09(input []string) {
	ints := []int{}
	for i, v := range input[0] {
		var d int
		if i%2 == 0 {
			d = i / 2
		} else {
			d = -1
		}
		for n := 0; n < int(v-'0'); n++ {
			ints = append(ints, d)
		}
	}
	temp := make([]int, len(ints))
	copy(temp, ints)
	fmt.Println(day9_part1(temp))
	fmt.Println(day9_part2(ints))
}

func day9_part1(arr []int) int {
	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {
		for arr[left] > -1 {
			left++
		}
		for arr[right] == -1 {
			right--
		}
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	return fileChecksum(arr)
}

func day9_part2(arr []int) int {
	for right := len(arr) - 1; right > 0; {
		for arr[right] == -1 {
			right--
		}
		rs := right
		for arr[right] == arr[rs] && right > 0 {
			right--
		}
		blocksize := rs - right
		for left := 0; left < right+1; {
			for arr[left] > -1 {
				left++
			}
			ls := left
			for arr[left] == -1 && left < right+1 {
				left++
			}
			if blocksize <= left-ls {
				for l, r := ls, rs; rs-r < blocksize; l, r = l+1, r-1 {
					arr[l], arr[r] = arr[r], arr[l]
				}
				break
			}
		}
	}
	return fileChecksum(arr)
}

func fileChecksum(arr []int) (res int) {
	for ind, i := range arr {
		if i > -1 {
			res += ind * i
		}
	}
	return
}
