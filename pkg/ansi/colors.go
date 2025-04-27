// Package ansi provides named ANSI escape sequences
// and primitive functions to generate such sequences.
package ansi

import (
	"fmt"
)

// ANSI Escape sequence to set foreground color (16 colors)
const (
	Black         EscapeSequence = pref + "[0;30m"
	Red           EscapeSequence = pref + "[0;31m"
	Green         EscapeSequence = pref + "[0;32m"
	Yellow        EscapeSequence = pref + "[0;33m"
	Blue          EscapeSequence = pref + "[0;34m"
	Magenta       EscapeSequence = pref + "[0;35m"
	Cyan          EscapeSequence = pref + "[0;36m"
	White         EscapeSequence = pref + "[0;37m"
	BrightBlack   EscapeSequence = pref + "[0;90m"
	BrightRed     EscapeSequence = pref + "[0;91m"
	BrightGreen   EscapeSequence = pref + "[0;92m"
	BrightYellow  EscapeSequence = pref + "[0;93m"
	BrightBlue    EscapeSequence = pref + "[0;94m"
	BrightMagenta EscapeSequence = pref + "[0;95m"
	BrightCyan    EscapeSequence = pref + "[0;96m"
	BrightWhite   EscapeSequence = pref + "[0;97m"
)

// ANSI Escape sequence to set background color (16 colors)
const (
	BlackBg         EscapeSequence = pref + "[40m"
	RedBg           EscapeSequence = pref + "[41m"
	GreenBg         EscapeSequence = pref + "[42m"
	YellowBg        EscapeSequence = pref + "[43m"
	BlueBg          EscapeSequence = pref + "[44m"
	MagentaBg       EscapeSequence = pref + "[45m"
	CyanBg          EscapeSequence = pref + "[46m"
	WhileBg         EscapeSequence = pref + "[47m"
	BrightBlackBg   EscapeSequence = pref + "[100m"
	BrightRedBg     EscapeSequence = pref + "[101m"
	BrightGreenBg   EscapeSequence = pref + "[102m"
	BrightYellowBg  EscapeSequence = pref + "[103m"
	BrightBlueBg    EscapeSequence = pref + "[104m"
	BrightMagentaBg EscapeSequence = pref + "[105m"
	BrightCyanBg    EscapeSequence = pref + "[106m"
	BrightWhiteBg   EscapeSequence = pref + "[107m"
)

// ANSI escape sequence that sets foreground color (256 colors)
const ColorTemplate256 = pref + "[38;5;%dm"

// ANSI escape sequence that sets background color (256 colors)
const BgColorTemplate256 = pref + "[48;5;%dm"

// Color256 returns an ANSI escape sequence that sets foreground color specified by color parameter.
func Color256(color byte) EscapeSequence {
	return EscapeSequence(fmt.Sprintf(ColorTemplate256, color))
}

// BgColor256 returns an ANSI escape sequence that sets background color specified by color parameter.
func BgColor256(color byte) EscapeSequence {
	return EscapeSequence(fmt.Sprintf(BgColorTemplate256, color))
}
