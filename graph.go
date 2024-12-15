package main

type Map struct {
	cells string
	width int
}

func (m Map) Width() int {
	return m.width
}

func (m Map) Height() int {
	return len(m.cells) / m.width
}

func (m Map) ToComplexGrid() map[complex64]rune {
	ret := make(map[complex64]rune, m.Height()*m.Width())
	for ind, v := range m.cells {
		ret[m.GetPointComplex(ind)] = v
	}
	return ret
}

func (m Map) GetPointComplex(index int) complex64 {
	return complex(float32(index%m.Width()), float32(index/m.Width()))
}

type Point2d struct {
	x, y int
}

func (m Map) GetPoint(index int) Point2d {
	return Point2d{x: index % m.Width(), y: index / m.Height()}
}

type rectangle struct {
	left  Point2d // Upper left
	right Point2d // Bottom right
}

func (p *Point2d) adjacent() <-chan Point2d {
	ch := make(chan Point2d)
	go func() {
		ch <- Point2d{p.x + 1, p.y + 1}
		ch <- Point2d{p.x + 1, p.y}
		ch <- Point2d{p.x + 1, p.y - 1}
		ch <- Point2d{p.x, p.y - 1}
		ch <- Point2d{p.x - 1, p.y - 1}
		ch <- Point2d{p.x - 1, p.y}
		ch <- Point2d{p.x - 1, p.y + 1}
		ch <- Point2d{p.x, p.y + 1}
		close(ch)
	}()
	return ch
}
