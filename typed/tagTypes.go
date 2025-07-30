package typed

type TagName string

const (
	DivTag      TagName = "div"
	SpanTag     TagName = "span"
	PTag        TagName = "p"
	ATag        TagName = "a"
	ImgTag      TagName = "img"
	InputTag    TagName = "input"
	ButtonTag   TagName = "button"
	FormTag     TagName = "form"
	LabelTag    TagName = "label"
	TableTag    TagName = "table"
	TrTag       TagName = "tr"
	TdTag       TagName = "td"
	ThTag       TagName = "th"
	UlTag       TagName = "ul"
	OlTag       TagName = "ol"
	LiTag       TagName = "li"
	H1Tag       TagName = "h1"
	H2Tag       TagName = "h2"
	H3Tag       TagName = "h3"
	H4Tag       TagName = "h4"
	H5Tag       TagName = "h5"
	H6Tag       TagName = "h6"
	TextareaTag TagName = "textarea"
	SelectTag   TagName = "select"
	OptionTag   TagName = "option"
	CanvasTag   TagName = "canvas"
	ScriptTag   TagName = "script"
	LinkTag     TagName = "link"
	StyleTag    TagName = "style"
)

func (t TagName) String() string {
	return string(t)
}
