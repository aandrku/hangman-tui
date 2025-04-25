package scenes

import (
	"hangman-tui/pkg/core"
	"hangman-tui/pkg/scenes/gameplay"
	"hangman-tui/pkg/scenes/menus/home"
)

func NewHomeMenu() core.Scene {
	return home.NewMenu()
}

func NewGameplay(attempts int) core.Scene {
	return gameplay.New(attempts)
}
