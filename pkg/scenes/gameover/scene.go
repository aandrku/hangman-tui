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
	height = 3
)

var (
	borderStyle  = ansi.Color256(23)
	headerStyle  = ansi.Color256(108)
	contentStyle = ansi.Color256(180)
	cueStyle     = ansi.Color256(60)
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
	str := fmt.Sprintf("hidden word: %s", s.state.Word)
	header := ""
	cuemsg := "Press any key to get back to the home menu"

	switch s.state.Status() {
	case state.GameWon:
		header = " You won! "
	case state.GameLost:
		header = " You lost:( "
	}

	width := max(len(str), len(header)) + 4
	col := centx - width/2
	row := centy - height/2

	draw.Box(row, col, height, width, borderStyle)
	draw.String(str, contentStyle, row+1, col+2)

	col = centx - len(header)/2
	draw.String(header, headerStyle, row, col)

	col = centx - len(cuemsg)/2
	draw.String(cuemsg, cueStyle, row+5, col)
}
