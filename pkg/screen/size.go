package screen

import (
	"golang.org/x/term"
	"os"
	"sync/atomic"
)

var (
	screenBuffer  [][]Cell // represents the screen
	currentWidth  int
	currentHeight int
	needsResize   atomic.Bool // signals main routine to resize the screenBuffer

)

// hardcoded limits for the terminal size
const (
	minWidth  = 80
	minHeight = 24
)

// GetSize returns current screen width and height
func GetSize() (int, int) {
	return currentWidth, currentHeight
}

// GetWidth returns current screen width
func GetWidth() int {
	return currentWidth
}

// GetHeight returns current screen height
func GetHeight() int {
	return currentWidth
}

// updateSize updates screen dimensions based on the current terminal size
func updateSize() {
	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		panic(err)
	}

	currentWidth = width
	currentHeight = height

	screenBuffer = make([][]Cell, 0, currentHeight)
	for range currentHeight {
		row := make([]Cell, currentWidth)
		screenBuffer = append(screenBuffer, row)
	}

}
