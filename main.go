package main

import (
	"log"
	"time"

	"github.com/gdamore/tcell"

	"snake/internal/board"
	"snake/internal/game"
	"snake/internal/input"
	"snake/internal/ui"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatal(err)
	}
	if err := screen.Init(); err != nil {
		log.Fatal(err)
	}
	defer screen.Fini()

	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack))
	screen.Clear()

	g := game.New(40, 20)
	renderer := ui.New(screen)

	handler := input.New(screen)
	handler.Poll()

	ticker := time.NewTicker(150 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case dir := <-handler.Events():
			switch dir {
			case input.Quit:
				return
			default:
				g.SetDirection(board.Direction(dir))
			}
		case <-ticker.C:
			g.Update()
			renderer.Draw(g)
		}
	}
}
