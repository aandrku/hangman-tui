package hud

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/screen/draw"
)

func NewLettersDisplay() *LettersDisplay {
	return &LettersDisplay{
		crossed:  make(map[rune]bool),
		reveiled: make(map[rune]bool),
	}

}

const (
	width  = 21
	height = 8
)

type LettersDisplay struct {
	crossed  map[rune]bool
	reveiled map[rune]bool
}

func (ld *LettersDisplay) Cross(letter rune) {
	ld.crossed[letter] = true
}

func (ld *LettersDisplay) Reveil(letter rune) {
	ld.reveiled[letter] = true
}

func (ld *LettersDisplay) Render() {
	w, _ := screen.GetSize()
	row := 0
	col := w - width - 1

	draw.Box(row, col, height, width, ansi.Cyan)

	draw.String(" Letters ", ansi.Red, row, col+5)

	curr := 'a'
	last := 'z'

	colS := col + 2
	colE := col + width - 2

	col = colS
	row = row + 2
	for curr <= last {
		style := ansi.White
		if _, ok := ld.crossed[curr]; ok {
			style = ansi.Red
		}
		if _, ok := ld.reveiled[curr]; ok {
			style = ansi.Green
		}
		screen.DrawChar(curr, style, row, col)
		screen.DrawChar(' ', ansi.EscapeSequence(""), row, col+1)
		col += 2
		if col >= colE {
			col = colS
			row += +2
		}
		curr++
	}

}
