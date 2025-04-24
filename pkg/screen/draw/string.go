package draw

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
)

func String(str string, style ansi.EscapeSequence, row, col int) {
	for _, v := range str {
		screen.DrawChar(v, style, row, col)
		col++
	}
}
