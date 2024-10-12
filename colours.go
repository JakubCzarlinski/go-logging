package logging

const reset = "\033[0m"

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
	return WrapInColour(msg, bgGreen)
}

func WhiteBg(msg string) string {
	return WrapInColour(msg, bgWhite)
}

func YellowBg(msg string) string {
	return WrapInColour(msg, bgYellow)
}

func RedBg(msg string) string {
	return WrapInColour(msg, bgRed)
}

func BlueBg(msg string) string {
	return WrapInColour(msg, bgBlue)
}

func MagentaBg(msg string) string {
	return WrapInColour(msg, bgMagenta)
}

func CyanBg(msg string) string {
	return WrapInColour(msg, bgCyan)
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
