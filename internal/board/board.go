package board

type Board struct {
	Width  int
	Height int
}

type Point struct {
	X, Y int
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func (b Board) OutOfBoard(x, y int) bool {
	return x < 0 || y < 0 || x >= b.Width || y >= b.Height
}
