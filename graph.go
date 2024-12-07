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

type point2d struct {
	x, y int
}

type rectangle struct {
	left  point2d // Upper left
	right point2d // Bottom right
}

func (p *point2d) adjacent() <-chan point2d {
	ch := make(chan point2d)
	go func() {
		ch <- point2d{p.x + 1, p.y + 1}
		ch <- point2d{p.x + 1, p.y}
		ch <- point2d{p.x + 1, p.y - 1}
		ch <- point2d{p.x, p.y - 1}
		ch <- point2d{p.x - 1, p.y - 1}
		ch <- point2d{p.x - 1, p.y}
		ch <- point2d{p.x - 1, p.y + 1}
		ch <- point2d{p.x, p.y + 1}
		close(ch)
	}()
	return ch
}
