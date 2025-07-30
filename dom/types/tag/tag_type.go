package tag

type TagName string

const (
	Div      TagName = "div"
	Span     TagName = "span"
	P        TagName = "p"
	A        TagName = "a"
	Img      TagName = "img"
	Input    TagName = "input"
	Button   TagName = "button"
	Form     TagName = "form"
	Label    TagName = "label"
	Table    TagName = "table"
	Tr       TagName = "tr"
	Td       TagName = "td"
	Th       TagName = "th"
	Ul       TagName = "ul"
	Ol       TagName = "ol"
	Li       TagName = "li"
	H1       TagName = "h1"
	H2       TagName = "h2"
	H3       TagName = "h3"
	H4       TagName = "h4"
	H5       TagName = "h5"
	H6       TagName = "h6"
	Textarea TagName = "textarea"
	Select   TagName = "select"
	Option   TagName = "option"
	Canvas   TagName = "canvas"
	Script   TagName = "script"
	Link     TagName = "link"
	Style    TagName = "style"
)

func (t TagName) String() string {
	return string(t)
}
