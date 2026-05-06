package food

import board "snake/internal/board"

type Food interface {
	Satiety() int
	Symbol() rune
	Position() board.Point
	SetPosition(board.Point)
}

type Apple struct {
	satiety  int
	position board.Point
}
type Banana struct {
	satiety  int
	position board.Point
}
type Grape struct {
	satiety  int
	position board.Point
}

func (a *Apple) Satiety() int {
	return a.satiety
}
func (b *Banana) Satiety() int {
	return b.satiety
}
func (g *Grape) Satiety() int {
	return g.satiety
}

func (a *Apple) Symbol() rune {
	return 'A'
}
func (b *Banana) Symbol() rune {
	return 'B'
}
func (g *Grape) Symbol() rune {
	return 'G'
}

func NewApple() Food {
	return &Apple{satiety: 1}
}
func NewBanana() Food {
	return &Banana{satiety: 2}
}
func NewGrape() Food {
	return &Grape{satiety: -1}
}

func (a *Apple) Position() board.Point {
	return a.position
}
func (a *Apple) SetPosition(p board.Point) {
	a.position = p
}
func (b *Banana) Position() board.Point {
	return b.position
}
func (b *Banana) SetPosition(p board.Point) {
	b.position = p
}
func (g *Grape) Position() board.Point {
	return g.position
}
func (g *Grape) SetPosition(p board.Point) {
	g.position = p
}
