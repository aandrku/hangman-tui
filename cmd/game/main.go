package main

import (
	"flag"
	"hangman-tui/pkg/boot"
	"hangman-tui/pkg/core"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/words"
)

func main() {
	boot.StartUp()
	defer boot.Shutdown("")

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
	c.Run()

}
