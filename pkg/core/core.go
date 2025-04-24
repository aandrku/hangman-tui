package core

import (
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
		}
	}
	return instance
}

type Core struct {
	currentScene Scene
}

func (c *Core) Run() {
	Start()
	for {
		start := time.Now()

		if key, ok := getKey(); ok {
			c.currentScene.ProcessKey(key)

		}

		screen.Clear()
		c.currentScene.Render()
		screen.Update()

		toSleep := frameTime - time.Since(start)
		if toSleep > 0 {
			time.Sleep(toSleep)
		}
	}
}

func (c *Core) SetScene(scene Scene) {
	c.currentScene = scene
}
