package ansi

import (
	"fmt"
)

// ANSI Escape sequence that moves cursor to a certain position (template)
const cursorPositionTemplate = pref + "[%d;%dH"

// ANSI escape sequence to clear screen
const ClearScreen EscapeSequence = pref + "[2J"

const (
	// ANSI escape sequence that hides cursor
	HideCursor EscapeSequence = pref + "[?25l"
	// ANSI escape sequence that shows cursor
	ShowCursor EscapeSequence = pref + "[?25h"
)

// CursorPosition return an ANSI escape sequence that moves cursor to (line, col) position.
func CursorPosition(line, col int) EscapeSequence {
	return EscapeSequence(fmt.Sprintf(cursorPositionTemplate, line, col))
}
