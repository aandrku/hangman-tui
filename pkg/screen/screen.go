package screen

import (
	"golang.org/x/term"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

var (
	screenBuffer  [][]cell // represents the screen
	currentWidth  int
	currentHeight int
	currentStyle  string // represents current drawing character style

	needsResize atomic.Bool // signals main routine to resize the screenBuffer

)

// hardcoded limits for the terminal size
const (
	minWidth  = 80
	minHeight = 24
)

func Init() {
	updateSize()
	go sigWinChHandler()
}

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

	screenBuffer = make([][]cell, 0, currentHeight)
	for range currentHeight {
		row := make([]cell, currentWidth)
		screenBuffer = append(screenBuffer, row)
	}

}

// sigWinChHandler listens for the SIGWINCH signal
// and tells the main goroutine to update the screen size.
func sigWinChHandler() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGWINCH)

	for {
		<-sigChan
		needsResize.Store(true)
	}
}
