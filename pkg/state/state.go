package state

import (
	"hangman-tui/pkg/entities/hud"
)

type State struct {
	// For Gameplay

	LettersDisplay  *hud.LettersDisplay
	WordDisplay     *hud.WordDisplay
	AttemptsDisplay *hud.AttemptsDisplay

	// For gameplay and home menu

	Word             string
	ProcessedLetters map[rune]bool
}

func (s *State) Restart(word string, attempts int) {
	s.Word = word
	s.ProcessedLetters = make(map[rune]bool)
	s.AttemptsDisplay = hud.NewAttemptsDisplay(attempts)

	s.LettersDisplay = hud.NewLettersDisplay()
	s.WordDisplay = hud.NewWordDisplay(word)
}
