package hud

import (
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

	draw.Box(row, col, height, width, borderStyle)

	draw.String(" Attempts Left ", headerStyle, row, col+2)

	content := strconv.Itoa(d.attemtps)
	draw.String(content, contentStyle, row+1, col+2)

}
