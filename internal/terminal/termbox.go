package terminal

import "github.com/nsf/termbox-go"

type TermboxScreen struct{}

func New() Screen {
	return &TermboxScreen{}
}

func (t *TermboxScreen) Init() error {
	return termbox.Init()
}

func (t *TermboxScreen) Close() {
	termbox.Close()
}

func (t *TermboxScreen) Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorBlack)
}

func (t *TermboxScreen) Flush() {
	termbox.Flush()
}

func (t *TermboxScreen) SetCell(x, y int, ch rune, fg Color, bg Color) {
	termbox.SetCell(x, y, ch, toTermboxColor(fg), toTermboxColor(bg))
}

func (t *TermboxScreen) Size() (int, int) {
	return termbox.Size()
}

func (t *TermboxScreen) SetOutputMode() {
	termbox.SetOutputMode(termbox.Output256)
}

func toTermboxColor(c Color) termbox.Attribute {
	switch c {
	case ColorBlack:
		return termbox.ColorBlack
	case ColorRed:
		return termbox.ColorRed
	case ColorGreen:
		return termbox.ColorGreen
	case ColorYellow:
		return termbox.ColorYellow
	case ColorBlue:
		return termbox.ColorBlue
	case ColorMagenta:
		return termbox.ColorMagenta
	case ColorCyan:
		return termbox.ColorCyan
	case ColorWhite:
		return termbox.ColorWhite
	default:
		return termbox.ColorDefault
	}
}

func (t *TermboxScreen) PollEvent() Event {
	ev := termbox.PollEvent()

	switch ev.Type {
	case termbox.EventKey:
		return translateKeyEvent(ev)
	case termbox.EventResize:
		return Event{Type: EventResize}
	default:
		return Event{Type: EventQuit}
	}
}

func translateKeyEvent(ev termbox.Event) Event {
	switch ev.Key {
	case termbox.KeyEsc:
		return Event{Type: EventKey, Key: KeyEsc}
	case termbox.KeyCtrlC:
		return Event{Type: EventKey, Key: KeyCtrlC}
	case termbox.KeyArrowLeft:
		return Event{Type: EventKey, Key: KeyArrowLeft}
	case termbox.KeyArrowRight:
		return Event{Type: EventKey, Key: KeyArrowRight}
	case termbox.KeyArrowUp:
		return Event{Type: EventKey, Key: KeyArrowUp}
	case termbox.KeyArrowDown:
		return Event{Type: EventKey, Key: KeyArrowDown}
	default:
		if ev.Ch != 0 {
			return Event{
				Type: EventKey,
				Key:  KeyRune,
				Rune: ev.Ch,
			}
		}
	}

	return Event{Type: EventKey, Key: KeyUnknown}
}
