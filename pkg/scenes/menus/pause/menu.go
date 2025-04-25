package pause

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/core"
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
func NewMenu(gameplay core.Scene) *menu.Menu {
	m := &menu.Menu{}

	m.SetHeader(" Home Menu ")
	m.SetWidth(width)
	m.SetHeight(height)

	m.SetHeaderStyle(headerStyle)
	m.SetBorderStyle(borderStyle)
	m.SetOptionStyle(optionStyle)

	m.AddOption(ReturnToGameOption{gameplay: gameplay})
	m.AddOption(BackToHomeMenuOption{})
	m.AddOption(ReturnToShellOption{})

	return m
}
