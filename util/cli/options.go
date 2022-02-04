package cli

// This block contains CLI options like font color, background color, etc.
const (
	RED    = "\033[31m"
	BLACK  = "\033[30m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	BLUE   = "\033[34m"
	PURPLE = "\033[35m"
	CYAN   = "\033[36m"
	WHITE  = "\033[37m"

	BG_BLACK  = "\033[40m"
	BG_RED    = "\033[41m"
	BG_GREEN  = "\033[42m"
	BG_YELLOW = "\033[43m"
	BG_BLUE   = "\033[44m"
	BG_PURPLE = "\033[45m"
	BG_CYAN   = "\033[46m"
	BG_WHITE  = "\033[47m"

	EOL        = "\n"
	TAB        = "\t"
	RESET      = "\033[0m"
	ITALIC     = "\033[3m"
	UNDERLINE  = "\033[4m"
	BLINK_SLOW = "\033[5m"
	BLINK_FAST = "\033[6m"
	REVERSE    = "\033[7m"
	HIDE       = "\033[8m"
	CROSS      = "\033[9m"
)
