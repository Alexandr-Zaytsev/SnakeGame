package snake

import (
	"snake/internal/board"
)

type Snake struct {
	Body      []board.Point
	direction board.Direction
	next      board.Direction
}

func New(start board.Point) *Snake {
	return &Snake{
		Body: []board.Point{
			start,
			{X: start.X - 1, Y: start.Y},
			{X: start.X - 2, Y: start.Y},
		},
		direction: board.Right,
		next:      board.Right,
	}
}

func (s *Snake) SetDirection(d board.Direction) {
	opposite := map[board.Direction]board.Direction{
		board.Up:    board.Down,
		board.Down:  board.Up,
		board.Left:  board.Right,
		board.Right: board.Left,
	}
	if d != opposite[s.direction] {
		s.next = d
	}
}

func (s *Snake) Move() {
	s.direction = s.next
	head := s.Head()

	newHead := board.Point{
		X: head.X + dx(s.direction),
		Y: head.Y + dy(s.direction),
	}

	s.Body = append([]board.Point{newHead}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) Grow() {
	tail := s.Body[len(s.Body)-1]
	s.Body = append(s.Body, tail)
}

func (s *Snake) Head() board.Point {
	return s.Body[0]
}

func (s *Snake) CollidesWithSelf() bool {
	head := s.Head()
	for _, p := range s.Body[1:] {
		if p == head {
			return true
		}
	}
	return false
}

func (s *Snake) Length() int {
	return len(s.Body)
}

func dx(d board.Direction) int {
	switch d {
	case board.Left:
		return -1
	case board.Right:
		return 1
	}
	return 0
}

func dy(d board.Direction) int {
	switch d {
	case board.Up:
		return -1
	case board.Down:
		return 1
	}
	return 0
}

func (s *Snake) Occupies(p board.Point) bool {
	for _, part := range s.Body {
		if part == p {
			return true
		}
	}
	return false
}

func (s *Snake) Shrink() {
	if len(s.Body) > 1 {
		s.Body = s.Body[:len(s.Body)-1]
	}
}
