package ansi

const (
	CURSOR_OFF = "\x1B[?25l"
	CURSOR_ON  = "\x1B[?25h"
	RESET      = "\x1B[0m"

	ERASE_LINE       = "\x1B[0m\x1B[0K"
	ERASE_SCRN_AFTER = "\x1B[0m\x1B[0J"

	UNDERLINE_ON  = "\x1B[4m"
	UNDERLINE_OFF = "\x1B[24m"

	// The alternate screen buffer (as used by vim, less, nano...) has no
	// scrollback, so terminals pass PgUp/PgDn through to the program.
	ALT_SCREEN_ON  = "\x1B[?1049h\x1B[H"
	ALT_SCREEN_OFF = "\x1B[?1049l"
)

var (
	YELLOW = "\x1B[0;33m"
)
