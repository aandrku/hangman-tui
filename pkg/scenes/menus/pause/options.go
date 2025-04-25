package pause

import (
	"hangman-tui/pkg/core"
	"hangman-tui/pkg/scenes"
)

const (
	optionLength        = 37
	optionCounterLength = 3
)

type ReturnToGameOption struct {
	gameplay core.Scene
}

func (r ReturnToGameOption) OnLeft()  {}
func (r ReturnToGameOption) OnRight() {}
func (r ReturnToGameOption) OnEnter() {
	c := core.GetInstance()
	c.SetScene(r.gameplay)
}
func (r ReturnToGameOption) String() string { return " Return to the Game" }

type BackToHomeMenuOption struct {
}

func (b BackToHomeMenuOption) OnLeft()  {}
func (b BackToHomeMenuOption) OnRight() {}
func (b BackToHomeMenuOption) OnEnter() {
	homeMenu := scenes.NewHomeMenu()
	c := core.GetInstance()
	c.SetScene(homeMenu)
}
func (b BackToHomeMenuOption) String() string { return " Back to Home Menu" }

type ReturnToShellOption struct {
}

func (r ReturnToShellOption) OnLeft()  {}
func (r ReturnToShellOption) OnRight() {}
func (po ReturnToShellOption) OnEnter() {
	core.Exit("")
}
func (r ReturnToShellOption) String() string { return " Return to Shell" }
