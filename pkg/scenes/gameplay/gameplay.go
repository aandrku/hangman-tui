package gameplay

import (
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/input"
	"hangman-tui/pkg/scenes/scene"
	"hangman-tui/pkg/state"
	"strings"
	"unicode"
)

func New(manager scene.Manager) *Gameplay {
	g := &Gameplay{
		manager: manager,
		state:   manager.State(),
	}

	return g
}

type Gameplay struct {
	manager scene.Manager
	state   *state.State
}

func (g *Gameplay) ProcessKey() {
	key, ok := input.GetKey()
	if !ok {
		return
	}
	if key.Key == keyboard.KeyEsc {
		g.manager.SetScene(scene.PauseMenu)
	}

	if unicode.IsLetter(key.Rune) {
		letter := unicode.ToLower(key.Rune)
		g.ProcessLetter(letter)
	}
}

func (g *Gameplay) ProcessLetter(letter rune) {
	if strings.ContainsRune(g.state.Word, letter) {
		g.state.LettersDisplay.Reveil(letter)
	} else {
		g.state.LettersDisplay.Cross(letter)
	}

}

func (g *Gameplay) Render() {
	g.state.LettersDisplay.Render()

}
