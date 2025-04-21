package main

import (
	"hangman-tui/pkg/screen"
	"time"
)

func main() {
	screen.Init()

	width, height := screen.GetSize()

	for r := range height {
		for c := range width {
			screen.Clear()
			screen.DrawChar('C', "\033[0;31m\033[42m", r, c)
			screen.Update()

			time.Sleep(time.Second / 60)
		}
	}

}
