package screen

import "hangman-tui/pkg/ansi"

// Cell represents an individual Cell on the screen
type Cell struct {
	Char  rune
	Style ansi.EscapeSequence
}
