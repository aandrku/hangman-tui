package gameover

import (
	"fmt"
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/input"
	"hangman-tui/pkg/scenes/scene"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/screen/draw"
	"hangman-tui/pkg/state"
)

const (
	width  = 25
	height = 3
)

func NewScene(manager scene.Manager) *Scene {
	return &Scene{
		manager: manager,
		state:   manager.State(),
	}

}

type Scene struct {
	manager scene.Manager
	state   *state.State
}

func (s *Scene) ProcessKey() {
	_ = input.GetKeyBlocking()
	s.manager.SetScene(scene.HomeMenu)
}

func (s *Scene) Render() {
	centx, centy := screen.GetCenter()

	col := centx - width/2
	row := centy - height/2 - 4

	draw.Box(row, col, height, width, ansi.Red)

	str := fmt.Sprintf("hidden word: %s", s.state.Word)
	draw.String(str, ansi.Blue, row+1, col+2)

	switch s.state.Status() {
	case state.GameWon:
		str = " You won! "
	case state.GameLost:
		str = " You lost:( "
	}

	col = centx - len(str)/2
	draw.String(str, ansi.Green, row, col)

	str = "Press any key to get back to the main menu"
	col = centx - len(str)/2
	draw.String(str, ansi.Yellow, row+5, col)
}
