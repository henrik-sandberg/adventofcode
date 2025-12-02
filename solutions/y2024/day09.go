package y2024

import (
	"adventofcode/solutions/shared"
)

func Day09(input []string) (solution shared.Solution[int, int]) {
	disk1 := []int{}
	for i, v := range input[0] {
		var d int
		if i%2 == 0 {
			d = i / 2
		} else {
			d = -1
		}
		for range int(v - '0') {
			disk1 = append(disk1, d)
		}
	}
	disk2 := make([]int, len(disk1))
	copy(disk2, disk1)
	day09_part1(disk1)
	day09_part2(disk2)
	checksum := func(arr []int) (res int) {
		for ind, i := range arr {
			if i > -1 {
				res += ind * i
			}
		}
		return
	}
	solution.Part1 = checksum(disk1)
	solution.Part2 = checksum(disk2)
	return
}

func day09_part1(arr []int) {
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
}

func day09_part2(arr []int) {
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
}
