package typed

type PropertyName string

const (
	InnerHTMLProp       PropertyName = "innerHTML"
	TextContentProp     PropertyName = "textContent"
	ValueProp           PropertyName = "value"
	CheckedProp         PropertyName = "checked"
	DisabledProp        PropertyName = "disabled"
	ClassNameProp       PropertyName = "className"
	IDProp              PropertyName = "id"
	HrefProp            PropertyName = "href"
	SrcProp             PropertyName = "src"
	AltProp             PropertyName = "alt"
	TitleProp           PropertyName = "title"
	NameProp            PropertyName = "name"
	TypeProp            PropertyName = "type"
	PlaceholderProp     PropertyName = "placeholder"
	TabIndexProp        PropertyName = "tabIndex"
	ReadOnlyProp        PropertyName = "readOnly"
	SelectedProp        PropertyName = "selected"
	DraggableProp       PropertyName = "draggable"
	ContentEditableProp PropertyName = "contentEditable"
	StyleProp           PropertyName = "style"
)

func (p PropertyName) String() string {
	return string(p)
}
