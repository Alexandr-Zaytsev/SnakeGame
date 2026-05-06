package input

import (
	"github.com/gdamore/tcell"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
	Quit
	Pause
	None
)

type Handler struct {
	screen tcell.Screen
	events chan Direction
}

func New(screen tcell.Screen) *Handler {
	return &Handler{
		screen: screen,
		events: make(chan Direction, 1),
	}
}

func (h *Handler) Poll() {
	go func() {
		for {
			ev := h.screen.PollEvent()
			switch e := ev.(type) {
			case *tcell.EventKey:
				h.events <- mapKey(e)
			}
		}
	}()
}

func (h *Handler) Events() <-chan Direction {
	return h.events
}

func mapKey(e *tcell.EventKey) Direction {
	switch e.Key() {
	case tcell.KeyUp:
		return Up
	case tcell.KeyDown:
		return Down
	case tcell.KeyLeft:
		return Left
	case tcell.KeyRight:
		return Right
	case tcell.KeyEsc:
		return Quit
	default:
		switch e.Rune() {
		case 'w':
			return Up
		case 's':
			return Down
		case 'a':
			return Left
		case 'd':
			return Right
		}
	}
	return None
}
