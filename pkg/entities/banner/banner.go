package banner

import (
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/screen/draw"
)

type BannerModel []string

var bannerModel = []string{
	"    __  __                                           ________  ______",
	"   / / / /___ _____  ____ _____ ___  ____ _____     /_  __/ / / /  _/",
	"  / /_/ / __ `/ __ \\/ __ `/ __ `__ \\/ __ `/ __ \\     / / / / / // /  ",
	" / __  / /_/ / / / / /_/ / / / / / / /_/ / / / /    / / / /_/ // /   ",
	"/_/ /_/\\__,_/_/ /_/\\__, /_/ /_/ /_/\\__,_/_/ /_/    /_/  \\____/___/   ",
	"                  /____/                                             ",
}

type Banner struct {
	model  BannerModel
	col    int
	row    int
	color1 ansi.EscapeSequence
	color2 ansi.EscapeSequence
}

func New() Banner {
	centerCol, centerRow := screen.GetCenter()

	col := centerCol - len(bannerModel[0])/2
	row := centerRow - len(bannerModel) - 3

	b := Banner{
		col:    col,
		row:    row,
		model:  bannerModel,
		color1: ansi.Color256(122),
		color2: ansi.Color256(180),
	}

	return b
}

func (b Banner) Render() {
	row := b.row
	col := b.col
	offset := 48
	for _, v := range b.model {
		half1 := v[:offset]
		half2 := v[offset:]
		draw.String(half1, b.color1, row, col)
		draw.String(half2, b.color2, row, col+offset)
		row++
	}
}
