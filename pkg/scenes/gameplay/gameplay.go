package gameplay

import (
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/input"
	"hangman-tui/pkg/scenes/scene"
)

func New(manager scene.Manager) *Gameplay {
	g := &Gameplay{manager: manager}

	return g
}

type Gameplay struct {
	manager scene.Manager
}

func (g *Gameplay) ProcessKey() {
	key, ok := input.GetKey()
	if !ok {
		return
	}
	if key.Key == keyboard.KeyEsc {
		g.manager.SetScene(scene.PauseMenu)
	}
}

func (g *Gameplay) Render() {

}
