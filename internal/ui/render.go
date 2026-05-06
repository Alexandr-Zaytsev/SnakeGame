package ui

import (
	"fmt"

	"snake/internal/board"
	"snake/internal/game"

	"github.com/gdamore/tcell"
)

var (
	styleDefault   = tcell.StyleDefault
	styleSnakeHead = tcell.StyleDefault.Foreground(tcell.ColorGreen).Bold(true)
	styleSnakeBody = tcell.StyleDefault.Foreground(tcell.ColorDarkGreen)
	styleBorder    = tcell.StyleDefault.Foreground(tcell.ColorWhite)
	styleFood      = tcell.StyleDefault.Foreground(tcell.ColorRed)
	styleScore     = tcell.StyleDefault.Foreground(tcell.ColorYellow)
	styleGameOver  = tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true)
)

type Render struct {
	screen tcell.Screen
}

func New(screen tcell.Screen) *Render {
	return &Render{screen: screen}
}

func (r *Render) Draw(g *game.Game) {
	r.screen.Clear()

	r.drawBorder(g.Board)
	r.drawFood(g)
	r.drawSnake(g)
	r.drawScore(g)

	if g.IsOver() {
		r.drawGameOver(g.Board)
	}

	r.screen.Show()
}

func (r *Render) drawBorder(b *board.Board) {
	for x := 0; x <= b.Width+1; x++ {
		r.set(x, 0, '─', styleBorder)
		r.set(x, b.Height+1, '─', styleBorder)
	}
	for y := 0; y <= b.Height+1; y++ {
		r.set(0, y, '│', styleBorder)
		r.set(b.Width+1, y, '│', styleBorder)
	}
	r.set(0, 0, '┌', styleBorder)
	r.set(b.Width+1, 0, '┐', styleBorder)
	r.set(0, b.Height+1, '└', styleBorder)
	r.set(b.Width+1, b.Height+1, '┘', styleBorder)
}

func (r *Render) drawSnake(g *game.Game) {
	for i, p := range g.Snake.Body {
		style := styleSnakeBody
		symbol := '■'
		if i == 0 {
			style = styleSnakeHead
			symbol = '●'
		}
		r.set(p.X+1, p.Y+1, symbol, style)
	}
}

func (r *Render) drawFood(g *game.Game) {
	p := g.Food.Position()
	r.set(p.X+1, p.Y+1, g.Food.Symbol(), styleFood)
}

func (r *Render) drawScore(g *game.Game) {
	score := fmt.Sprintf(" Score: %d ", g.Score)
	for i, ch := range score {
		r.set(i+2, 0, ch, styleScore)
	}
}

func (r *Render) drawGameOver(b *board.Board) {
	r.drawCentered(b, "GAME OVER", styleGameOver, 0)
	r.drawCentered(b, "press ESC to exit", styleGameOver, 1)
}

func (r *Render) drawCentered(b *board.Board, msg string, style tcell.Style, offset int) {
	x := (b.Width-len(msg))/2 + 1
	y := b.Height/2 + 1 + offset
	for i, ch := range msg {
		r.set(x+i, y, ch, style)
	}
}

func (r *Render) set(x, y int, ch rune, style tcell.Style) {
	r.screen.SetContent(x, y, ch, nil, style)
}
