package home

import (
	"fmt"
	"hangman-tui/pkg/boot"
	"hangman-tui/pkg/scenes/scene"
	"hangman-tui/pkg/words"
	"strconv"
)

const (
	optionLength        = 37
	optionCounterLength = 3
)

type PlayOption struct {
	manager scene.Manager
}

func (po PlayOption) OnLeft()  {}
func (po PlayOption) OnRight() {}
func (po PlayOption) OnEnter() {
	state := po.manager.State()
	state.Restart("hello", attemptsCounter)
	po.manager.SetScene(scene.Gameplay)
}
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
	l := s.CurrentLength()
	lStr := fmt.Sprintf("<%*s > ", optionCounterLength, strconv.Itoa(l))
	return fmt.Sprintf("%-*s %s", optionLength-optionCounterLength-3, " Word Length", lStr)
}

type AttemptsCountOption struct {
}

func (ao AttemptsCountOption) OnLeft() {
	decreaseAttempts()

}
func (ao AttemptsCountOption) OnRight() {
	increaseAttempts()
}
func (ao AttemptsCountOption) OnEnter() {}
func (ao AttemptsCountOption) String() string {
	a := attemptsCounter
	aStr := fmt.Sprintf("<%*s > ", optionCounterLength, strconv.Itoa(a))
	return fmt.Sprintf("%-*s %s", optionLength-optionCounterLength-3, " Attempts", aStr)
}

type ExitOption struct {
}

func (eo ExitOption) OnLeft()  {}
func (eo ExitOption) OnRight() {}
func (eo ExitOption) OnEnter() {
	boot.Shutdown("")
}
func (eo ExitOption) String() string { return " Exit" }
