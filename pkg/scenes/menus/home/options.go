package home

import (
	"fmt"
	"hangman-tui/pkg/core"
	"hangman-tui/pkg/words"
	"strconv"
)

const (
	optionLength        = 37
	optionCounterLength = 3
)

type PlayOption struct {
}

func (po PlayOption) OnLeft()        {}
func (po PlayOption) OnRight()       {}
func (po PlayOption) OnEnter()       {}
func (po PlayOption) String() string { return " Play" }

type WordLengthOption struct {
}

func (wo WordLengthOption) OnLeft() {
	s := words.GetStore()
	s.DecreaseLength()
}
func (wo WordLengthOption) OnRight() {
	s := words.GetStore()
	s.IncreaseLength()
}
func (wo WordLengthOption) OnEnter() {}
func (wo WordLengthOption) String() string {
	s := words.GetStore()
	length := s.CurrentLength()
	lengthStr := fmt.Sprintf("<%*s >", optionCounterLength, strconv.Itoa(length))
	return fmt.Sprintf("%-*s %s", optionLength-optionCounterLength-2, " Word Length", lengthStr)
}

type AttemptsCountOption struct {
}

func (ao AttemptsCountOption) OnLeft()        {}
func (ao AttemptsCountOption) OnRight()       {}
func (ao AttemptsCountOption) OnEnter()       {}
func (ao AttemptsCountOption) String() string { return " Attempts" }

type ExitOption struct {
}

func (eo ExitOption) OnLeft()  {}
func (eo ExitOption) OnRight() {}
func (eo ExitOption) OnEnter() {
	core.Exit("")
}
func (eo ExitOption) String() string { return " Exit" }
