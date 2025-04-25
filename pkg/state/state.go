package state

import (
	"hangman-tui/pkg/entities/hud"
)

type StateStatus int

const (
	GameWon = iota
	GameLost
	GameRunning
)

type State struct {
	// For Gameplay

	LettersDisplay  *hud.LettersDisplay
	WordDisplay     *hud.WordDisplay
	AttemptsDisplay *hud.AttemptsDisplay

	// For gameplay and home menu

	Word             string
	ProcessedLetters map[rune]bool

	// For gameplay and game summary
	status StateStatus
}

func (s *State) Status() StateStatus {
	return s.status
}

func (s *State) SetStatus(status StateStatus) {
	s.status = status
}

func (s *State) Restart(word string, attempts int) {
	s.Word = word
	s.ProcessedLetters = make(map[rune]bool)
	s.AttemptsDisplay = hud.NewAttemptsDisplay(attempts)

	s.LettersDisplay = hud.NewLettersDisplay()
	s.WordDisplay = hud.NewWordDisplay(word)

	s.status = GameRunning
}
