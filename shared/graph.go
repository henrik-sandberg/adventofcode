package shared

type Grid map[complex64]byte

func (g *Grid) FindAny(b byte) complex64 {
	for k, v := range *g {
		if v == b {
			return k
		}
	}
	return 0
}

func (g *Grid) FindAll(b byte) []complex64 {
	res := []complex64{}
	for k, v := range *g {
		if v == b {
			res = append(res, k)
		}
	}
	return res
}

func (g *Grid) Get(x, y int) byte {
	return (*g)[complex(float32(x), float32(y))]
}

func NewGrid(arr []string) Grid {
	ret := make(map[complex64]byte, len(arr)*len(arr[0]))
	for ri, row := range arr {
		for ci, byt := range []byte(row) {
			ret[complex(float32(ci), float32(ri))] = byt
		}
	}
	return ret
}

func (g *Grid) Up() complex64 {
	return -1i
}

func (g *Grid) Down() complex64 {
	return 1i
}

func (g *Grid) Left() complex64 {
	return -1
}

func (g *Grid) Right() complex64 {
	return 1
}

func (g *Grid) Directions() []complex64 {
	return []complex64{g.Up(), g.Right(), g.Down(), g.Left()}
}
