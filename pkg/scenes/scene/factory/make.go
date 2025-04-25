package factory

import (
	"hangman-tui/pkg/scenes/gameover"
	"hangman-tui/pkg/scenes/gameplay"
	"hangman-tui/pkg/scenes/menus/home"
	"hangman-tui/pkg/scenes/menus/pause"
	"hangman-tui/pkg/scenes/scene"
)

func Make(manager scene.Manager, id scene.ID) scene.Scene {
	switch id {
	case scene.HomeMenu:
		return home.NewMenu(manager)
	case scene.Gameplay:
		return gameplay.New(manager)
	case scene.PauseMenu:
		return pause.NewMenu(manager)
	case scene.GameOver:
		return gameover.NewScene(manager)
	}

	panic("Scene factory should always return from a switch statement")
}
