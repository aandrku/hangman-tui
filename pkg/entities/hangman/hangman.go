package hangman

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
)

const height = 8
const width = 12
const StagesCount = 10

var stages = []func(*[height][width]screen.Cell){
	drawPost,
	drawTopBar,
	drawRope,
	drawHead,
	drawUpperTorse,
	drawLowerTorso,
	drawLeftArm,
	drawRightArm,
	drawLeftLeg,
	drawRightLeg,
}

func New(stage int) *Hangman {
	h := &Hangman{}

	empty := screen.Cell{
		Char:  ' ',
		Style: ansi.EscapeSequence(""),
	}

	buff := [height][width]screen.Cell{}
	for r := range buff {
		for c := range buff[0] {
			buff[r][c] = empty
		}
	}

	drawBase(&buff)
	for s := range stage {
		stages[s](&buff)
	}
	h.buff = buff
	h.stage = stage

	return h
}

type Hangman struct {
	buff  [height][width]screen.Cell
	stage int
}

func (h *Hangman) NextStage() {
	if h.stage == len(stages) {
		panic("stage should never grow beyond limits")
	}
	stages[h.stage](&h.buff)
	h.stage++
}

func (h *Hangman) Render() {
	screenWidth, screenHeight := screen.GetSize()
	screenCenterCol, _ := screen.GetCenter()
	rowIdx := screenHeight - len(h.buff) - 2
	colStart := screenCenterCol - 10
	colIdx := colStart

	for c := range screenWidth {
		screen.DrawChar('═', groundStyle, screenHeight-3, c)
	}

	for r := range height {
		for c := range width {
			screen.DrawCell(h.buff[r][c], rowIdx, colIdx)
			colIdx++
		}
		colIdx = colStart
		rowIdx++
	}
}
