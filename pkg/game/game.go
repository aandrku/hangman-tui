package game

import (
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/boot"
	"hangman-tui/pkg/scenes/scene"
	"hangman-tui/pkg/scenes/scene/factory"
	"hangman-tui/pkg/screen"
	"sync"
	"time"
)

var lock sync.Mutex
var g *Game

const frameTime = time.Second / 60

func New() *Game {
	g = &Game{}
	g.SetScene(scene.HomeMenu)

	return g
}

type Game struct {
	scene scene.Scene
}

func (g *Game) Run() {
	// open keyboard
	var err error
	keyboardInput, err = keyboard.GetKeys(1)
	if err != nil {
		boot.Shutdown("Hangman TUI failed to open keyboard: " + err.Error())
	}

	for {
		start := time.Now()

		if key, ok := getKey(); ok {
			g.scene.ProcessKey(key)

		}

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
