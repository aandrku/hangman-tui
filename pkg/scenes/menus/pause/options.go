package pause

import (
	"hangman-tui/pkg/boot"
	"hangman-tui/pkg/scenes/scene"
)

const (
	optionLength        = 37
	optionCounterLength = 3
)

type ReturnToGameOption struct {
	manager scene.Manager
}

func (r ReturnToGameOption) OnLeft()  {}
func (r ReturnToGameOption) OnRight() {}
func (r ReturnToGameOption) OnEnter() {
	r.manager.SetScene(scene.Gameplay)
}
func (r ReturnToGameOption) String() string { return " Return to the Game" }

type BackToHomeMenuOption struct {
	manager scene.Manager
}

func (b BackToHomeMenuOption) OnLeft()  {}
func (b BackToHomeMenuOption) OnRight() {}
func (b BackToHomeMenuOption) OnEnter() {
	b.manager.SetScene(scene.HomeMenu)
}
func (b BackToHomeMenuOption) String() string { return " Back to Home Menu" }

type ReturnToShellOption struct {
}

func (r ReturnToShellOption) OnLeft()  {}
func (r ReturnToShellOption) OnRight() {}
func (po ReturnToShellOption) OnEnter() {
	boot.Exit("")
}
func (r ReturnToShellOption) String() string { return " Return to Shell" }
