package state

import (
	"hangman-tui/pkg/entities/hud"
)

type State struct {
	// For Gameplay

	LettersDisplay *hud.LettersDisplay
	WordDisplay    *hud.WordDisplay

	// For gameplay and home menu

	Word             string
	Attempts         int
	ProcessedLetters map[rune]bool
}

func (s *State) Restart(word string, attempts int) {
	s.Word = word
	s.Attempts = attempts
	s.ProcessedLetters = make(map[rune]bool)

	s.LettersDisplay = hud.NewLettersDisplay()
	s.WordDisplay = hud.NewWordDisplay(word)
}
