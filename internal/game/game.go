package game

import (
	"math/rand"

	"snake/internal/board"
	"snake/internal/food"
	"snake/internal/snake"
)

type State int

const (
	Running State = iota
	Paused
	Over
)

type Game struct {
	Board *board.Board
	Snake *snake.Snake
	Food  food.Food
	Score int
	State State
}

func New(width, height int) *Game {
	b := &board.Board{Width: width, Height: height}

	start := board.Point{X: width / 2, Y: height / 2}
	s := snake.New(start)

	g := &Game{
		Board: b,
		Snake: s,
		State: Running,
	}
	g.spawnFood()
	return g
}

func (g *Game) Update() {
	if g.State != Running {
		return
	}

	g.Snake.Move()

	head := g.Snake.Head()

	head.X = (head.X + g.Board.Width) % g.Board.Width
	head.Y = (head.Y + g.Board.Height) % g.Board.Height
	g.Snake.Body[0] = head

	if g.Snake.CollidesWithSelf() {
		g.State = Over
		return
	}

	if g.Board.OutOfBoard(head.X, head.Y) {
		g.State = Over
		return
	}

	if head == g.Food.Position() {
		if g.Food.Satiety() < 0 {
			g.Snake.Shrink()
		} else {
			g.Snake.Grow()
		}
		g.Score += g.Food.Satiety()
		g.spawnFood()
	}
}

func (g *Game) SetDirection(d board.Direction) {
	if g.State == Running {
		g.Snake.SetDirection(d)
	}
}

func (g *Game) IsOver() bool {
	return g.State == Over
}

func (g *Game) spawnFood() {
	foods := []food.Food{
		food.NewApple(),
		food.NewBanana(),
		food.NewGrape(),
	}

	for {
		p := board.Point{
			X: rand.Intn(g.Board.Width),
			Y: rand.Intn(g.Board.Height),
		}
		if !g.Snake.Occupies(p) {
			f := foods[rand.Intn(len(foods))]
			f.SetPosition(p)
			g.Food = f
			return
		}
	}
}
