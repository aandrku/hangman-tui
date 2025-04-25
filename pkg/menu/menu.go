package menu

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/screen/draw"
)

type Menu struct {
	currentOption int
	options       []MenuOption
	header        string

	width  int
	height int

	headerStyle ansi.EscapeSequence
	borderStyle ansi.EscapeSequence
	optionStyle ansi.EscapeSequence
}

func (m *Menu) ProcessKey(key keyboard.KeyEvent) {
	index := m.currentOption
	switch {
	case key.Rune == 'j' || key.Key == keyboard.KeyArrowDown:
		m.NextOption()
	case key.Rune == 'k' || key.Key == keyboard.KeyArrowUp:
		m.PrevOption()
	case key.Rune == 'h' || key.Key == keyboard.KeyArrowLeft:
		m.options[index].OnLeft()
	case key.Rune == 'l' || key.Key == keyboard.KeyArrowRight:
		m.options[index].OnRight()
	case key.Key == keyboard.KeyEnter:
		m.options[index].OnEnter()
	}
}

func (m *Menu) Render() {
	centerCol, centerRow := screen.GetCenter()
	startRow := centerRow - m.height/2
	startCol := centerCol - m.width/2

	headerStartCol := centerCol - len(m.header)/2

	// draw border
	draw.Box(startRow, startCol, m.height, m.width, m.borderStyle)

	// draw header
	draw.String(m.header, m.headerStyle, startRow, headerStartCol)

	//draw options
	optionRow := startRow + 1
	optionCol := startCol + 1
	for i, v := range m.options {
		var style ansi.EscapeSequence
		if i == m.currentOption {
			style = m.optionStyle + ansi.Reverse
		} else {
			style = m.optionStyle
		}
		str := fmt.Sprintf("%-*s", m.width-1, v.String())
		draw.String(str, style, optionRow, optionCol)
		optionRow++
	}

}

func (m *Menu) AddOption(option MenuOption) {
	m.options = append(m.options, option)
}

func (m *Menu) NextOption() {
	m.currentOption++
	if m.currentOption >= len(m.options) {
		m.currentOption = 0
	}

}

func (m *Menu) PrevOption() {
	m.currentOption--
	if m.currentOption < 0 {
		m.currentOption = len(m.options) - 1
	}

}

func (m *Menu) SetHeader(header string) {
	m.header = header
}

func (m *Menu) SetWidth(width int) {
	m.width = width
}

func (m *Menu) SetHeight(height int) {
	m.height = height
}

func (m *Menu) SetHeaderStyle(style ansi.EscapeSequence) {
	m.headerStyle = style
}

func (m *Menu) SetBorderStyle(style ansi.EscapeSequence) {
	m.borderStyle = style
}

func (m *Menu) SetOptionStyle(style ansi.EscapeSequence) {
	m.optionStyle = style
}

type MenuOption interface {
	String() string
	OnLeft()
	OnRight()
	OnEnter()
}
