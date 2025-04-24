package core

import (
	"github.com/eiannone/keyboard"
)

func getKey() (keyboard.KeyEvent, bool) {
	select {
	case keyEvent := <-keyboardInput:
		preprocessKey(keyEvent)
		return keyEvent, true
	default:
		return keyboard.KeyEvent{}, false
	}
}

func preprocessKey(key keyboard.KeyEvent) {
	if key.Key == keyboard.KeyCtrlC {
		Exit("")
	}
}
