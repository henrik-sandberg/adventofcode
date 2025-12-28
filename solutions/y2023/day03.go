package y2023

import (
	"adventofcode/solutions/shared"
	"regexp"
	"strconv"
)

func Day03(input []string) (solution shared.Solution[int, int]) {
	re := regexp.MustCompile(`\d+`)
	grid := shared.NewGrid(input)
	numbers := make([][]rectangle, len(input))
	for ri, row := range input {
		for _, match := range re.FindAllStringIndex(row, -1) {
			numbers[ri] = append(numbers[ri], rectangle{
				point2d{match[0], ri},
				point2d{match[1], ri},
			})
		}
	}
	for point, char := range grid {
		if char == '.' || char >= '0' && char <= '9' {
			continue
		}
		sni := getSurroundingNumberIndexes(point, grid, numbers)
		sn := make([]int, 0, len(sni))
		for ni := range sni {
			start := ni.left
			end := ni.right
			number, _ := strconv.Atoi(input[start.y][start.x:end.x])
			sn = append(sn, number)
		}
		solution.Part1 += shared.Sum(sn...)
		if len(sn) == 2 {
			solution.Part2 += shared.Product(sn...)
		}
	}
	return
}

func getSurroundingNumberIndexes(point complex128, grid shared.Grid, numbers [][]rectangle) map[rectangle]bool {
	ret := make(map[rectangle]bool)
	for _, adj := range []complex128{
		1, -1, 1i, -1i, 1 + 1i, 1 - 1i, -1 + 1i, -1 - 1i,
	} {
		p := point + adj
		if grid[p] < '0' || grid[p] > '9' {
			continue
		}
		x, y := int(real(p)), int(imag(p))
		for _, numberIndexes := range numbers[y] {
			start := numberIndexes.left
			end := numberIndexes.right
			if start.x <= x && x <= end.x {
				ret[rectangle{left: start, right: end}] = true
			}
		}
	}
	return ret
}
