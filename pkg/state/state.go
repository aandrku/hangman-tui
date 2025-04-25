package state

import (
	"hangman-tui/pkg/entities/hud"
)

type State struct {
	LettersDisplay *hud.LettersDisplay
	Word           string
	Attempts       int
}

func (s *State) Restart(word string, attempts int) {
	s.Word = word
	s.Attempts = attempts
	s.LettersDisplay = hud.NewLettersDisplay()
}
