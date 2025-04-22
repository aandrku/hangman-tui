package screen

import (
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	updateSize()
	go sigWinChHandler()
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
