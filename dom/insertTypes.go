package dom

type InsertionMethod string

const (
	AppendChildMethod  InsertionMethod = "appendChild"
	PrependMethod      InsertionMethod = "prepend"
	InsertBeforeMethod InsertionMethod = "insertBefore"
	ReplaceChildMethod InsertionMethod = "replaceChild"
	RemoveChildMethod  InsertionMethod = "removeChild"
	RemoveMethod       InsertionMethod = "remove"
	AfterMethod        InsertionMethod = "after"
	BeforeMethod       InsertionMethod = "before"
	ReplaceWithMethod  InsertionMethod = "replaceWith"
)

func (p InsertionMethod) String() string {
	return string(p)
}
