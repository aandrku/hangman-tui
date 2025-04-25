package boot

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/ansi"
	"os"
)

func Exit(msg string) {
	// show cursor
	fmt.Print(ansi.ShowCursor)
	fmt.Print(ansi.ClearScreen)

	fmt.Print(ansi.CursorPosition(1, 1))
	fmt.Print(ansi.Red)
	if msg == "" {
		fmt.Print("\nThank you for playing", ansi.Magenta, " Hangman TUI\n\n")
	} else {
		fmt.Println(msg)
	}
	fmt.Print(ansi.Reset)

	// close keyboard
	err := keyboard.Close()
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
