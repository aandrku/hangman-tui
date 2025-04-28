// Package boot handles the startup and shutdown routines for the Hangman TUI game.
//
// It provides functions to initialize the game environment (such as loading custom words,
// setting up input handling, and preparing the terminal display) and to cleanly terminate
// the game session by resetting the terminal and closing resources.
package boot

import (
	"flag"
	"fmt"
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/input"
	"hangman-tui/pkg/screen"
	"hangman-tui/pkg/words"
)

// StartUp initializes the Hangman game environment.
//
// It hides the cursor in the terminal, parses the command-line flag "-words" to load a custom word list if provided,
// initializes the input handling system and screen rendering system.
//
// If any initialization step fails, the program will shut down gracefully
// with an appropriate error message.
func StartUp() {
	// hide cursor
	fmt.Print(ansi.HideCursor)

	// load custom words if specified by user
	wordsFile := flag.String("words", "", "usage: hangman -words<path/to/the/file/with/words>")
	flag.Parse()
	if *wordsFile != "" {
		s := words.GetStore()
		err := s.ReadFromFile(*wordsFile)
		if err != nil {
			Shutdown(err.Error())
		}
	}

	err := input.Init()
	if err != nil {
		Shutdown(err.Error())
	}

	// initialize screen package
	screen.Init()
}
