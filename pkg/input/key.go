package input

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"os"
)

var keyboardInput <-chan keyboard.KeyEvent

func Init() error {
	var err error
	keyboardInput, err = keyboard.GetKeys(1)
	if err != nil {
		return fmt.Errorf("could not open keyboard: %v", err)
	}
	return nil
}

func Close() {
	_ = keyboard.Close
}

func preprocessKey(key keyboard.KeyEvent) {
	if key.Key == keyboard.KeyCtrlC {
		os.Exit(0)
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

func GetKeyBlocking() keyboard.KeyEvent {
	keyEvent := <-keyboardInput
	preprocessKey(keyEvent)
	return keyEvent
}
