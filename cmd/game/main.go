package main

import (
	"hangman-tui/pkg/boot"
	"hangman-tui/pkg/game"
)

func main() {
	boot.StartUp()

	// start the core with home menu
	c := game.New()
	c.Run()

}
