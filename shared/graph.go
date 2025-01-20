package shared

type Grid map[complex128]byte

func (g *Grid) FindAny(b byte) complex128 {
	for k, v := range *g {
		if v == b {
			return k
		}
	}
	return 0
}

func (g *Grid) FindAll(b byte) []complex128 {
	res := []complex128{}
	for k, v := range *g {
		if v == b {
			res = append(res, k)
		}
	}
	return res
}

func NewGrid(arr []string) Grid {
	ret := make(map[complex128]byte, len(arr)*len(arr[0]))
	for ri, row := range arr {
		for ci, byt := range []byte(row) {
			ret[complex(float64(ci), float64(ri))] = byt
		}
	}
	return ret
}

func (g *Grid) Up() complex128 {
	return -1i
}

func (g *Grid) Down() complex128 {
	return 1i
}

func (g *Grid) Left() complex128 {
	return -1
}

func (g *Grid) Right() complex128 {
	return 1
}

func (g *Grid) Directions() []complex128 {
	return []complex128{g.Up(), g.Right(), g.Down(), g.Left()}
}
