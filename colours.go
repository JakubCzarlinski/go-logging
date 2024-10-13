package logging

const reset = "\033[0m"
const resetBg = "\033[49m"

const (
	bgGreen   = "\033[97;42m"
	bgWhite   = "\033[90;47m"
	bgYellow  = "\033[90;43m"
	bgRed     = "\033[97;41m"
	bgBlue    = "\033[97;44m"
	bgMagenta = "\033[97;45m"
	bgCyan    = "\033[97;46m"
)
const (
	green   = "\033[32m"
	white   = "\033[97m"
	yellow  = "\033[33m"
	red     = "\033[31m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func GreenBg(msg string) string {
	return WrapInColourBg(msg, bgGreen)
}

func WhiteBg(msg string) string {
	return WrapInColourBg(msg, bgWhite)
}

func YellowBg(msg string) string {
	return WrapInColourBg(msg, bgYellow)
}

func RedBg(msg string) string {
	return WrapInColourBg(msg, bgRed)
}

func BlueBg(msg string) string {
	return WrapInColourBg(msg, bgBlue)
}

func MagentaBg(msg string) string {
	return WrapInColourBg(msg, bgMagenta)
}

func CyanBg(msg string) string {
	return WrapInColourBg(msg, bgCyan)
}

func Green(msg string) string {
	return WrapInColour(msg, green)
}

func White(msg string) string {
	return WrapInColour(msg, white)
}

func Yellow(msg string) string {
	return WrapInColour(msg, yellow)
}

func Red(msg string) string {
	return WrapInColour(msg, red)
}

func Blue(msg string) string {
	return WrapInColour(msg, blue)
}

func Magenta(msg string) string {
	return WrapInColour(msg, magenta)
}

func Cyan(msg string) string {
	return WrapInColour(msg, cyan)
}

func WrapInColour(msg string, colour string) string {
	return colour + msg + reset
}

func WrapInColourBg(msg string, colourBg string) string {
	return colourBg + msg + resetBg
}

func Bold(msg string) string {
	return "\033[1m" + msg + "\033[22m"
}

func Dim(msg string) string {
	return "\033[2m" + msg + "\033[22m"
}

func Italic(msg string) string {
	return "\033[3m" + msg + "\033[23m"
}

func Underline(msg string) string {
	return "\033[4m" + msg + "\033[24m"
}

func Blink(msg string) string {
	return "\033[5m" + msg + "\033[25m"
}

func FastBlink(msg string) string {
	return "\033[6m" + msg + "\033[25m"
}

func Reverse(msg string) string {
	return "\033[7m" + msg + "\033[27m"
}

func Hidden(msg string) string {
	return "\033[8m" + msg + "\033[28m"
}

func StrikeThrough(msg string) string {
	return "\033[9m" + msg + "\033[29m"
}

func DoubleUnderline(msg string) string {
	return "\033[21m" + msg + "\033[24m"
}

func Overlined(msg string) string {
	return "\033[53m" + msg + "\033[55m"
}

func Combine(formats ...ColourFunc) ColourFunc {
	return func(msg string) string {
		for _, format := range formats {
			msg = format(msg)
		}
		return msg
	}
}
