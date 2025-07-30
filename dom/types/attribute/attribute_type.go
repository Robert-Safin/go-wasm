package attribute

type AttributeName string

const (
	InnerHTML       AttributeName = "innerHTML"
	TextContent     AttributeName = "textContent"
	Value           AttributeName = "value"
	Checked         AttributeName = "checked"
	DisabledA       AttributeName = "disabled"
	ClassName       AttributeName = "className"
	ID              AttributeName = "id"
	Href            AttributeName = "href"
	Src             AttributeName = "src"
	Alt             AttributeName = "alt"
	Title           AttributeName = "title"
	Name            AttributeName = "name"
	Type            AttributeName = "type"
	Placeholder     AttributeName = "placeholder"
	TabIndex        AttributeName = "tabIndex"
	ReadOnly        AttributeName = "readOnly"
	Selected        AttributeName = "selected"
	Draggable       AttributeName = "draggable"
	ContentEditable AttributeName = "contentEditable"
	Style           AttributeName = "style"
)

func (p AttributeName) String() string {
	return string(p)
}
