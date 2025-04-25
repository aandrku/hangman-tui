package core

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
var instance *Core

const frameTime = time.Second / 60

func GetInstance() *Core {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Core{}
			instance.SetScene(scene.HomeMenu)
		}
	}
	return instance
}

type Core struct {
	scene scene.Scene
}

func (c *Core) Run() {
	// open keyboard
	var err error
	keyboardInput, err = keyboard.GetKeys(1)
	if err != nil {
		boot.Shutdown("Hangman TUI failed to open keyboard: " + err.Error())
	}

	for {
		start := time.Now()

		if key, ok := getKey(); ok {
			c.scene.ProcessKey(key)

		}

		screen.Clear()
		c.scene.Render()
		screen.Update()

		toSleep := frameTime - time.Since(start)
		if toSleep > 0 {
			time.Sleep(toSleep)
		}
	}
}

func (c *Core) SetScene(sceneId scene.ID) {
	c.scene = factory.Make(c, sceneId)
}
