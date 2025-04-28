package boot

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"hangman-tui/pkg/ansi"
	"hangman-tui/pkg/input"
	"os"
)

// Shutdown gracefully terminates the Hangman game.
//
// It shows the cursor, clears the screen, displays a farewell message or an error message in red,
// resets the terminal colors, closes the keyboard input and exits the program.
//
// If closing the keyboard input fails, the function panics.
func Shutdown(msg string) {
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
	if err := keyboard.Close(); err != nil {
		panic(err)
	}

	input.Close()

	os.Exit(0)
}
