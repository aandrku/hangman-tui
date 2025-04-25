package game

import (
	"hangman-tui/pkg/scenes/scene"
	"hangman-tui/pkg/scenes/scene/factory"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/state"
	"time"
)

var g *Game

const frameTime = time.Second / 60

func New() *Game {
	g = &Game{
		state: &state.State{},
	}
	g.SetScene(scene.HomeMenu)

	return g
}

type Game struct {
	state *state.State
	scene scene.Scene
}

func (g *Game) Run() {

	screen.Clear()
	g.scene.Render()
	screen.Update()

	for {
		start := time.Now()
		g.scene.ProcessKey()

		screen.Clear()
		g.scene.Render()
		screen.Update()

		toSleep := frameTime - time.Since(start)
		if toSleep > 0 {
			time.Sleep(toSleep)
		}
	}
}

func (g *Game) SetScene(sceneId scene.ID) {
	g.scene = factory.Make(g, sceneId)
}

func (g *Game) State() *state.State {
	return g.state
}
