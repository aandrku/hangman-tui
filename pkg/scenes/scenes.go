package scenes

import (
	"hangman-tui/pkg/menu"
	"hangman-tui/pkg/scenes/menus/home"
)

func NewHomeMenu() *menu.Menu {
	return home.NewMenu()
}
