package home

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/menu"
	"hangman-tui/pkg/scenes/scene"
)

var (
	borderStyle = ansi.Color256(23)
	headerStyle = ansi.Color256(108)
	optionStyle = ansi.Color256(180)
)

const (
	width  = 40
	height = 7
)

// Menu represents a home menu scene.
func NewMenu(manager scene.Manager) *menu.Menu {
	m := menu.New()

	m.SetHeader(" Home Menu ")
	m.SetWidth(width)
	m.SetHeight(height)

	m.SetHeaderStyle(headerStyle)
	m.SetBorderStyle(borderStyle)
	m.SetOptionStyle(optionStyle)

	m.AddOption(PlayOption{manager: manager})
	m.AddOption(WordLengthOption{})
	m.AddOption(AttemptsCountOption{})
	m.AddOption(ExitOption{})

	return m
}
