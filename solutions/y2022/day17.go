package y2022

import (
	"adventofcode/solutions/shared"
)

func Day17(input []string) (solution shared.Solution[int, int]) {
	var shapes = [][]int{
		{0b0011110},
		{0b0001000, 0b0011100, 0b0001000},
		{0b0011100, 0b0000100, 0b0000100},
		{0b0010000, 0b0010000, 0b0010000, 0b0010000},
		{0b0011000, 0b0011000},
	}
	const (
		width     = 7
		leftWall  = 1 << (width - 1)
		rightWall = 1
	)
	canShift := func(tower []int, block []int, y int, dir byte) bool {
		for i, row := range block {
			if dir == '<' && row&leftWall != 0 || dir == '>' && row&rightWall != 0 {
				return false
			}
			if y+i < len(tower) &&
				(tower[y+i]&row != 0 ||
					dir == '<' && tower[y+i]&(row<<1) != 0 ||
					dir == '>' && tower[y+i]&(row>>1) != 0) {
				return false
			}
		}
		return true
	}
	type CacheKey struct {
		shapeIdx int
		jetIdx   int
	}
	type CacheVal struct {
		step   int
		height int
	}
	jets := input[0]
	solve := func(target int) int {
		memo := make(map[CacheKey]CacheVal)
		jetIdx := 0
		shapeIdx := 0
		var tower []int
		for step := range target {
			key := CacheKey{
				shapeIdx: shapeIdx,
				jetIdx:   jetIdx,
			}
			if cache, ok := memo[key]; ok {
				cycleLen := step - cache.step
				cycleHeight := len(tower) - cache.height

				remaining := target - step
				q, extra := remaining/cycleLen, remaining%cycleLen
				if extra == 0 {
					return len(tower) + q*cycleHeight
				}
			}
			memo[key] = CacheVal{
				step:   step,
				height: len(tower),
			}
			orig := shapes[shapeIdx]
			block := make([]int, len(orig))
			copy(block, orig)
			shapeIdx = (shapeIdx + 1) % len(shapes)
			y := len(tower) + 3
			for {
				dir := jets[jetIdx]
				if canShift(tower, block, y, dir) {
					switch dir {
					case '>':
						for idx := range block {
							block[idx] >>= 1
						}
					case '<':
						for idx := range block {
							block[idx] <<= 1
						}
					}
				}
				jetIdx = (jetIdx + 1) % len(jets)
				if y == 0 || !canShift(tower, block, y-1, 0) {
					break
				}
				y--
			}
			for i, row := range block {
				if y+i >= len(tower) {
					tower = append(tower, row)
				} else {
					tower[y+i] |= row
				}
			}
		}
		return len(tower)
	}
	solution.Part1 = solve(2022)
	solution.Part2 = solve(1e12)
	return
}
