package scene

import (
	"hangman-tui/pkg/state"
)

type Manager interface {
	SetScene(sceneId ID)
	State() *state.State
}
