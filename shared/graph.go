package shared

type Grid map[complex64]rune

func (g *Grid) FindAny(r rune) complex64 {
	for k, v := range *g {
		if v == r {
			return k
		}
	}
	return 0
}

func (g *Grid) FindAll(r rune) []complex64 {
	res := []complex64{}
	for k, v := range *g {
		if v == r {
			res = append(res, k)
		}
	}
	return res
}

func (g *Grid) Get(x, y int) rune {
	return (*g)[complex(float32(x), float32(y))]
}

func NewGrid(arr []string) Grid {
	ret := make(map[complex64]rune, len(arr)*len(arr[0]))
	for ri, row := range arr {
		for ci, cell := range row {
			ret[complex(float32(ci), float32(ri))] = cell
		}
	}
	return ret
}
