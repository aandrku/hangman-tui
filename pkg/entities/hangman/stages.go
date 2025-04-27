package hangman

import (
	"hangman-tui/pkg/screen"
)

func drawBase(buff *[height][width]screen.Cell) {
	const baseStartRow = height - 1
	const baseStartCol = 0
	const baseRune = '▀'

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
	row := height - 2
	col := 1
	postRune1 := '║'

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
		Style: topBarStyle,
	}
	col++

	for ; col < 10; col++ {
		buff[row][col] = screen.Cell{
			Char:  rune2,
			Style: topBarStyle,
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
		Style: topBarStyle,
	}
	row++
	buff[row][col] = screen.Cell{
		Char:  rune2,
		Style: ropeStyle,
	}

}
func drawHead(buff *[height][width]screen.Cell) {
	row := 2
	col := 10
	rune := 'O'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: headStyle,
	}

}
func drawTorso(buff *[height][width]screen.Cell) {
	row := 3
	col := 10
	rune := '|'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: lowerTorsoStyle,
	}

}
func drawLeftArm(buff *[height][width]screen.Cell) {
	row := 3
	col := 9
	rune := '/'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: armsStyle,
	}

}
func drawRightArm(buff *[height][width]screen.Cell) {
	row := 3
	col := 11
	rune := '\\'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: armsStyle,
	}

}
func drawLeftLeg(buff *[height][width]screen.Cell) {
	row := 4
	col := 9
	rune := '/'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: legsStyle,
	}

}
func drawRightLeg(buff *[height][width]screen.Cell) {
	row := 4
	col := 11
	rune := '\\'

	buff[row][col] = screen.Cell{
		Char:  rune,
		Style: legsStyle,
	}

}
