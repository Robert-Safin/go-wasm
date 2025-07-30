package attribute

type AttributeName string

const (
	InnerHTMLAttribute       AttributeName = "innerHTML"
	TextContentAttribute     AttributeName = "textContent"
	ValueAttribute           AttributeName = "value"
	CheckedAttribute         AttributeName = "checked"
	DisabledAttribute        AttributeName = "disabled"
	ClassNameAttribute       AttributeName = "className"
	IDAttribute              AttributeName = "id"
	HrefAttribute            AttributeName = "href"
	SrcAttribute             AttributeName = "src"
	AltAttribute             AttributeName = "alt"
	TitleAttribute           AttributeName = "title"
	NameAttribute            AttributeName = "name"
	TypeAttribute            AttributeName = "type"
	PlaceholderAttribute     AttributeName = "placeholder"
	TabIndexAttribute        AttributeName = "tabIndex"
	ReadOnlyAttribute        AttributeName = "readOnly"
	SelectedAttribute        AttributeName = "selected"
	DraggableAttribute       AttributeName = "draggable"
	ContentEditableAttribute AttributeName = "contentEditable"
	StyleAttribute           AttributeName = "style"
)

func (p AttributeName) String() string {
	return string(p)
}
