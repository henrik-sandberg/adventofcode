package y2023

type point2d struct {
	x, y int
}

type rectangle struct {
	left  point2d // Upper left
	right point2d // Bottom right
}
