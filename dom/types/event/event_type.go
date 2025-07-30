package event

type EventType string

const (
	ClickEvent       EventType = "click"
	DblClickEvent    EventType = "dblclick"
	InputEvent       EventType = "input"
	ChangeEvent      EventType = "change"
	SubmitEvent      EventType = "submit"
	FocusEvent       EventType = "focus"
	BlurEvent        EventType = "blur"
	KeyDownEvent     EventType = "keydown"
	KeyUpEvent       EventType = "keyup"
	KeyPressEvent    EventType = "keypress"
	MouseDownEvent   EventType = "mousedown"
	MouseUpEvent     EventType = "mouseup"
	MouseMoveEvent   EventType = "mousemove"
	MouseEnterEvent  EventType = "mouseenter"
	MouseLeaveEvent  EventType = "mouseleave"
	ScrollEvent      EventType = "scroll"
	WheelEvent       EventType = "wheel"
	TouchStartEvent  EventType = "touchstart"
	TouchEndEvent    EventType = "touchend"
	TouchMoveEvent   EventType = "touchmove"
	ResizeEvent      EventType = "resize"
	ContextMenuEvent EventType = "contextmenu"
	LoadEvent        EventType = "load"
	ErrorEvent       EventType = "error"
)

func (e EventType) String() string {
	return string(e)
}
