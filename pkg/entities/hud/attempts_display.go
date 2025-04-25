package hud

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen/draw"
	"strconv"
)

func NewAttemptsDisplay(attempts int) *AttemptsDisplay {

	return &AttemptsDisplay{
		attemtps: attempts,
	}
}

type AttemptsDisplay struct {
	attemtps int
}

func (d *AttemptsDisplay) Decrement() int {
	d.attemtps--
	return d.attemtps
}

func (d *AttemptsDisplay) Render() {
	row := 0
	col := 0
	width := 20
	height := 3

	draw.Box(row, col, height, width, ansi.Cyan)

	draw.String(" Attempts Left ", ansi.Red, row, col+2)

	content := strconv.Itoa(d.attemtps)
	draw.String(content, ansi.Magenta, row+1, col+2)

}
