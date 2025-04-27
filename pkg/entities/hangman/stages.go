package hangman

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
)

func drawBase(buff *[height][width]screen.Cell) {
	const baseStartRow = height - 1
	const baseStartCol = 0
	const baseStyle = ansi.Red
	const baseRune = '═'

	row := baseStartRow
	col := baseStartCol

	for ; col < 12; col++ {
		buff[row][col] = screen.Cell{
			Char:  baseRune,
			Style: baseStyle,
		}

	}
}

func drawPost(buff *[height][width]screen.Cell) {
	postStartRow := height - 1
	postStartCol := 1
	postStyle := ansi.Green
	postRune1 := '║'
	postRune2 := '╩'

	row, col := postStartRow, postStartCol
	buff[row][col] = screen.Cell{
		Char:  postRune2,
		Style: postStyle,
	}
	row--

	for ; row > 0; row-- {
		buff[row][col] = screen.Cell{
			Char:  postRune1,
			Style: postStyle,
		}

	}
}

func drawTopBar(buff *[height][width]screen.Cell) {
	row := 0
	col := 1
	rune1 := '╦'
	rune2 := '═'

	buff[row][col] = screen.Cell{
		Char:  rune1,
		Style: ansi.Yellow,
	}
	col++

	for ; col < 10; col++ {
		buff[row][col] = screen.Cell{
			Char:  rune2,
			Style: ansi.Yellow,
		}

	}

}
func drawRope(buff *[height][width]screen.Cell) {
	row := 0
	col := 10
	rune1 := '╦'
	rune2 := '║'

	buff[row][col] = screen.Cell{
		Char:  rune1,
		Style: ansi.Yellow,
	}
	row++
	buff[row][col] = screen.Cell{
		Char:  rune2,
		Style: ansi.Blue,
	}

}
func drawHead(buff *[height][width]screen.Cell) {
	row := 2
	col := 10
	rune := 'O'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Red,
	}

}
func drawUpperTorse(buff *[height][width]screen.Cell) {
	row := 3
	col := 10
	rune := '|'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Magenta,
	}
}
func drawLowerTorso(buff *[height][width]screen.Cell) {
	row := 4
	col := 10
	rune := '|'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Magenta,
	}

}
func drawLeftArm(buff *[height][width]screen.Cell) {
	row := 3
	col := 9
	rune := '/'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Magenta,
	}

}
func drawRightArm(buff *[height][width]screen.Cell) {
	row := 3
	col := 11
	rune := '\\'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Magenta,
	}

}
func drawLeftLeg(buff *[height][width]screen.Cell) {
	row := 5
	col := 9
	rune := '/'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Magenta,
	}

}
func drawRightLeg(buff *[height][width]screen.Cell) {
	row := 5
	col := 11
	rune := '\\'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: ansi.Magenta,
	}

}
