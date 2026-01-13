package terminal

type Color uint8

const (
	ColorDefault Color = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

type Screen interface {
	Init() error
	Close()
	Clear()
	Flush()
	SetOutputMode()

	SetCell(x, y int, ch rune, fg, bg Color)
	Size() (width, height int)

	PollEvent() Event
}
