package draw

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
)

// Light horizontal and vertical lines
const (
	LightHorizontal rune = '─' // ─
	LightVertical   rune = '│' // │
)

// Light corners and intersections
const (
	LightTopLeft     rune = '┌' // ┌
	LightTopRight    rune = '┐' // ┐
	LightBottomLeft  rune = '└' // └
	LightBottomRight rune = '┘' // ┘
	LightTeeUp       rune = '┬' // ┬
	LightTeeDown     rune = '┴' // ┴
	LightTeeLeft     rune = '├' // ├
	LightTeeRight    rune = '┤' // ┤
	LightCross       rune = '┼' // ┼
)

func Box(row, col, height, width int, style ansi.EscapeSequence) {
	rowIndex := row
	lastRow := row + height - 1
	colIdex := col
	lastCol := col + width - 1

	//draw vertical lines
	for range height - 1 {
		screen.DrawChar(LightVertical, style, rowIndex, col)
		screen.DrawChar(LightVertical, style, rowIndex, lastCol)
		rowIndex++
	}

	// draw horizontal lines
	for range width - 1 {
		screen.DrawChar(LightHorizontal, style, row, colIdex)
		screen.DrawChar(LightHorizontal, style, lastRow, colIdex)
		colIdex++
	}

	// draw corners
	screen.DrawChar(LightTopLeft, style, row, col)
	screen.DrawChar(LightTopRight, style, row, lastCol)
	screen.DrawChar(LightBottomLeft, style, lastRow, col)
	screen.DrawChar(LightBottomRight, style, lastRow, lastCol)

}
