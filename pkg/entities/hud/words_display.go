package hud

import (
	"hangman-tui/pkg/screen/draw"
)

func NewWordDisplay(word string) *WordDisplay {
	display := ""
	for range word {
		display += "_"
	}

	return &WordDisplay{
		word:    []rune(word),
		display: []rune(display),
	}

}

type WordDisplay struct {
	word    []rune
	display []rune
}

func (d *WordDisplay) Reveil(letter rune) {
	for i, v := range d.word {
		if v == letter {
			d.display[i] = letter
		}
	}
}

func (d *WordDisplay) IsGuessed() bool {
	return string(d.word) == string(d.display)
}

func (d *WordDisplay) Render() {
	row := 0
	col := 22

	width := 20
	height := 3

	content := string(d.display)

	draw.Box(row, col, height, width, borderStyle)

	draw.String(" Word ", headerStyle, row, col+2)
	draw.String(content, contentStyle, row+1, col+2)

}
