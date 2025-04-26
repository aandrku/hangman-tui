package main

import (
	"flag"
	"hangman-tui/pkg/boot"
	"hangman-tui/pkg/game"
	"hangman-tui/pkg/input"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/words"
)

func main() {
	boot.StartUp()

	input.Init()
	defer input.Close()

	// load custom words if specified by user
	wordsFile := flag.String("words", "", "usage: hangman -words<path/to/the/file/with/words>")
	flag.Parse()
	if *wordsFile != "" {
		s := words.GetStore()
		s.ReadFromFile(*wordsFile)
	}

	// initialize screen package
	screen.Init()

	// start the core with home menu
	c := game.New()
	c.Run()

}
