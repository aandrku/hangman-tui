package scene

import (
	"github.com/eiannone/keyboard"
)

type Scene interface {
	ProcessKey(key keyboard.KeyEvent)
	Render()
}
