package gameplay

import (
	"github.com/eiannone/keyboard"
)

func New(attempts int) *Gameplay {
	g := &Gameplay{}

	return g
}

type Gameplay struct {
}

func (g *Gameplay) ProcessKey(key keyboard.KeyEvent) {

}

func (g *Gameplay) Render() {

}
