package terminal

type EventType int

const (
	EventKey EventType = iota
	EventResize
	EventQuit
)

type Key int

const (
	KeyUnknown Key = iota
	KeyRune
	KeyEsc
	KeyCtrlC
	KeyArrowLeft
	KeyArrowRight
	KeyArrowUp
	KeyArrowDown
)

type Event struct {
	Type EventType
	Key  Key
	Rune rune
}
