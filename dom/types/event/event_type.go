package event

type EventType string

const (
	Click       EventType = "click"
	DblClick    EventType = "dblclick"
	Input       EventType = "input"
	Change      EventType = "change"
	Submit      EventType = "submit"
	Focus       EventType = "focus"
	Blur        EventType = "blur"
	KeyDown     EventType = "keydown"
	KeyUp       EventType = "keyup"
	KeyPress    EventType = "keypress"
	MouseDown   EventType = "mousedown"
	MouseUp     EventType = "mouseup"
	MouseMove   EventType = "mousemove"
	MouseEnter  EventType = "mouseenter"
	MouseLeave  EventType = "mouseleave"
	Scroll      EventType = "scroll"
	Wheel       EventType = "wheel"
	TouchStart  EventType = "touchstart"
	TouchEnd    EventType = "touchend"
	TouchMove   EventType = "touchmove"
	Resize      EventType = "resize"
	ContextMenu EventType = "contextmenu"
	Load        EventType = "load"
	Error       EventType = "error"
)

func (e EventType) String() string {
	return string(e)
}
