package scene

import ()

type Scene interface {
	ProcessKey()
	Render()
}
