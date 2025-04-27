package screen

import (
	"fmt"
	"hangman-tui/pkg/ansi"
	"strings"
)

// Clear clears the screenBuffer.
func Clear() {
	emptyCell := Cell{
		Char:  ' ',
		Style: "",
	} // used to clean the screen

	for r := range len(screenBuffer) {
		for c := range len(screenBuffer[0]) {
			screenBuffer[r][c] = emptyCell
		}
	}
}

// Update displays the contents of the screenBuffer in the terminal.
func Update() {
	if currentHeight < minHeight || currentWidth < minWidth {
		drawSmallScreenWarning()
	}

	builder := strings.Builder{}
	currStyle := ansi.EscapeSequence("")

	for row := range currentHeight {
		for col := range currentWidth {
			char := screenBuffer[row][col].Char
			style := screenBuffer[row][col].Style
			cursor := fmt.Sprintf("\033[%d;%dH", row+1, col+1)

			if currStyle != style {
				currStyle = style
				builder.WriteString("\033[0m")
				builder.WriteString(string(style))
			}

			builder.WriteString(cursor)
			builder.WriteRune(char)

		}
	}

	fmt.Print(builder.String())

	// handle terminal size change asynchronously
	if resize := needsResize.Swap(false); resize {
		updateSize()
	}
}

// drawSmallScreenWarning replaces screenBuffer content with warning.
func drawSmallScreenWarning() {

}

// DrawChar draws a character on the screen.
func DrawChar(char rune, style ansi.EscapeSequence, row, column int) {
	// guards
	if row < 0 || row >= currentHeight {
		return
	}
	if column < 0 || column >= currentWidth {
		return
	}

	c := Cell{
		Char:  char,
		Style: style,
	}
	screenBuffer[row][column] = c
}

// DrawCell draws a cell on the screen.
func DrawCell(cell Cell, row, column int) {
	// guards
	if row < 0 || row >= currentHeight {
		return
	}
	if column < 0 || column >= currentWidth {
		return
	}

	screenBuffer[row][column] = cell
}
