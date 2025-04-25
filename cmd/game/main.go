package main

import (
	"flag"
	"hangman-tui/pkg/core"
	"hangman-tui/pkg/scenes/menus/home"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/words"
)

func main() {
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
	c := core.GetInstance()
	c.SetScene(home.NewMenu())
	c.Run()

}
