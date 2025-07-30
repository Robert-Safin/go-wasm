package insert

type InsertionMethod string

const (
	AppendChild  InsertionMethod = "appendChild"
	Prepend      InsertionMethod = "prepend"
	InsertBefore InsertionMethod = "insertBefore"
	ReplaceChild InsertionMethod = "replaceChild"
	RemoveChild  InsertionMethod = "removeChild"
	Remove       InsertionMethod = "remove"
	After        InsertionMethod = "after"
	Before       InsertionMethod = "before"
	ReplaceWith  InsertionMethod = "replaceWith"
)

func (p InsertionMethod) String() string {
	return string(p)
}
