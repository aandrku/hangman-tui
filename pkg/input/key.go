package input

import (
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/boot"
)

var keyboardInput <-chan keyboard.KeyEvent

func Init() {
	var err error
	keyboardInput, err = keyboard.GetKeys(1)
	if err != nil {
		boot.Shutdown("Hangman TUI failed to open keyboard: " + err.Error())
	}
}

func Close() {
	_ = keyboard.Close
}

func preprocessKey(key keyboard.KeyEvent) {
	if key.Key == keyboard.KeyCtrlC {
		boot.Shutdown("")
	}
}

func GetKey() (keyboard.KeyEvent, bool) {
	select {
	case keyEvent := <-keyboardInput:
		preprocessKey(keyEvent)
		return keyEvent, true
	default:
		return keyboard.KeyEvent{}, false
	}
}
