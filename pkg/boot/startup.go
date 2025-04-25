package boot

import (
	"fmt"
	"hangman-tui/pkg/ansi"
)

func StartUp() {
	// hide cursor
	fmt.Print(ansi.HideCursor)

	// open keyboard
}
