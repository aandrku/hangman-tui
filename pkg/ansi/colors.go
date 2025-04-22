package ansi

import (
	"fmt"
)

// ANSI Escape sequence to set foreground color (16 colors)
const (
	Black         = pref + "[0;30m"
	Red           = pref + "[0;31m"
	Green         = pref + "[0;32m"
	Yellow        = pref + "[0;33m"
	Blue          = pref + "[0;34m"
	Magenta       = pref + "[0;35m"
	Cyan          = pref + "[0;36m"
	White         = pref + "[0;37m"
	BrightBlack   = pref + "[0;90m"
	BrightRed     = pref + "[0;91m"
	BrightGreen   = pref + "[0;92m"
	BrightYellow  = pref + "[0;93m"
	BrightBlue    = pref + "[0;94m"
	BrightMagenta = pref + "[0;95m"
	BrightCyan    = pref + "[0;96m"
	BrightWhite   = pref + "[0;97m"
)

// ANSI Escape sequence to set background color (16 colors)
const (
	BlackBg         = pref + "[40m"
	RedBg           = pref + "[41m"
	GreenBg         = pref + "[42m"
	YellowBg        = pref + "[43m"
	BlueBg          = pref + "[44m"
	MagentaBg       = pref + "[45m"
	CyanBg          = pref + "[46m"
	WhileBg         = pref + "[47m"
	BrightBlackBg   = pref + "[100m"
	BrightRedBg     = pref + "[101m"
	BrightGreenBg   = pref + "[102m"
	BrightYellowBg  = pref + "[103m"
	BrightBlueBg    = pref + "[104m"
	BrightMagentaBg = pref + "[105m"
	BrightCyanBg    = pref + "[106m"
	BrightWhiteBg   = pref + "[107m"
)

// ANSI escape sequence that sets foreground color (256 colors)
const ColorTemplate256 = pref + "[38;5;%d"

// ANSI escape sequence that sets background color (256 colors)
const BgColorTemplate256 = pref + "[48;5;%d"

// Color256 returns an ANSI escape sequence that sets foreground color specified by color parameter.
func Color256(color byte) string {
	return fmt.Sprintf(ColorTemplate256, color)
}

// BgColor256 returns an ANSI escape sequence that sets background color specified by color parameter.
func BgColor256(color byte) string {
	return fmt.Sprintf(BgColorTemplate256, color)
}
