package pause

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/menu"
	"hangman-tui/pkg/scenes/scene"
)

const (
	headerStyle = ansi.Green
	borderStyle = ansi.Red
	optionStyle = ansi.Cyan
)

const (
	width  = 40
	height = 7
)

// Menu represents a home menu scene.
func NewMenu(manager scene.Manager) *menu.Menu {
	m := &menu.Menu{}

	m.SetHeader(" Game Paused ")
	m.SetWidth(width)
	m.SetHeight(height)

	m.SetHeaderStyle(headerStyle)
	m.SetBorderStyle(borderStyle)
	m.SetOptionStyle(optionStyle)

	m.AddOption(ReturnToGameOption{manager: manager})
	m.AddOption(BackToHomeMenuOption{manager: manager})
	m.AddOption(ReturnToShellOption{})

	return m
}
