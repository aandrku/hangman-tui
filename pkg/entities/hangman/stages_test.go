package hangman

import (
	"hangman-tui/pkg/screen"
	"testing"
)

func TestDrawBase(t *testing.T) {
	buff := [height][width]screen.Cell{}

	t.Log(buff)

	drawBase(&buff)
	t.Log("after call")
	for _, v := range buff {
		t.Log(v, "\n")
	}
	t.Errorf("")

	drawPost(&buff)
	t.Log("after call")
	for _, v := range buff {
		t.Log(v, "\n")
	}
	t.Errorf("")

}
