package ansi

// Reset all modes (style and colors)
const Reset EscapeSequence = pref + "[0m"

// Set reverse mode
const Reverse EscapeSequence = pref + "[7m"

// Set crossed mode (crossed symbols)
const Crossed EscapeSequence = pref + "[9m"
