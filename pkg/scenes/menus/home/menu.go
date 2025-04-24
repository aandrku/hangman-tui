package home

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/menu"
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
func NewMenu() *menu.Menu {
	m := &menu.Menu{}

	m.SetHeader(" Home Menu ")
	m.SetWidth(width)
	m.SetHeight(height)

	m.SetHeaderStyle(headerStyle)
	m.SetBorderStyle(borderStyle)
	m.SetOptionStyle(optionStyle)

	m.AddOption(PlayOption{})
	m.AddOption(WordLengthOption{})
	m.AddOption(AttemptsCountOption{})
	m.AddOption(ExitOption{})

	return m
}
