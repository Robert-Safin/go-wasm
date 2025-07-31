//go:build js && wasm

package dom

import (
	"strings"
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

// Wrapper around js.Value for method implementation purposes.
type HtmlElement struct {
	// Value - Untyped access as escape catch.
	Value js.Value
}

// GETTERS AND SETTER

// Retuns string behind given attribute. Will return a false if attribute is not set.
// Note: browsers can and will accept objects as attribute values, by calling .toString() and returning that.
func (h HtmlElement) GetAttribute(prop attribute.AttributeName) (string, bool) {
	v := h.Value.Get(string(prop))
	if !v.Truthy() {
		var zero string
		return zero, false
	}
	return v.String(), true
}

// Sets attribute value.
func (h HtmlElement) SetAttribute(prop attribute.AttributeName, value string) {
	h.Value.Set(prop.String(), value)
}

// Sets multiple attribute values. (loops and calls SetAttribute)
func (h HtmlElement) SetAttributeMap(props map[attribute.AttributeName]string) {
	for k, v := range props {
		h.SetAttribute(k, v)
	}
}

// STYLING ATTRIBUTE

// Sets value of the style attribute. Accepts a map of property:value.
func (e HtmlElement) SetStyles(styles map[string]string) {
	joined := ""
	for k, v := range styles {
		joined += k + ":" + v + ";"
	}
	e.SetAttribute(attribute.Style, joined)
}

// Overlays and sets new styles on top of existing styles inside style attribute.
// If a property already existed, its value will be updated.
// If a property did not exist, it will be set.
func (e HtmlElement) UpdateStyles(newStyles map[string]string) {
	v := e.Value.Call("getAttribute", "style")
	existingStyles := v.String()

	split := strings.Split(existingStyles, ";")
	split = split[:len(split)-1]

	updatedStyles := make(map[string]string, len(split))

	for _, pair := range split {
		pair_split := strings.Split(pair, ":")
		prop := strings.TrimSpace(pair_split[0])
		value := strings.TrimSpace(pair_split[1])
		updatedStyles[prop] = value
	}
	for k, v := range newStyles {
		updatedStyles[k] = v
	}

	e.SetStyles(updatedStyles)
}

// Accepts properties, removes those properties and their values from style attribute.
func (e HtmlElement) RemoveStyles(styles ...string) {
	v := e.Value.Call("getAttribute", "style")
	existingStyles := v.String()

	split := strings.Split(existingStyles, ";")
	split = split[:len(split)-1]

	updatedStyles := make(map[string]string, len(split))
	for _, pair := range split {
		pair_split := strings.Split(pair, ":")
		prop := strings.TrimSpace(pair_split[0])
		value := strings.TrimSpace(pair_split[1])
		updatedStyles[prop] = value
	}
	for _, s := range styles {
		delete(updatedStyles, s)
	}

	e.SetStyles(updatedStyles)
}

// CLASSES

// Adds classes to class attribute.
func (e HtmlElement) AddClasses(classNames ...string) {
	js := e.Value.Get("classList")
	for _, v := range classNames {
		js.Call("add", v)
	}
}

// Remove classes from class attribute.
func (e HtmlElement) RemoveClasses(classNames ...string) {
	js := e.Value.Get("classList")
	for _, v := range classNames {
		js.Call("remove", v)
	}
}

// Check if class exists.
func (e HtmlElement) ContainsClass(className string) bool {
	js := e.Value.Get("classList")
	js = js.Call("contains", className)
	return js.Bool()
}

// Removes a class if it exists. If it does not exist, it is added.
func (e HtmlElement) ToggleClasses(classNames ...string) {
	js := e.Value.Get("classList")
	for _, className := range classNames {
		js = js.Call("toggle", className)
	}
}

// MISC

// Deletes element from DOM.
func (e HtmlElement) Delete() {
	e.Value.Call("remove")
}

// INSERTING

// Appends new child element as the last child of the target.
func (e HtmlElement) AppendChild(child HtmlElement) {
	e.Value.Call("append", child.Value)
}

// Appends new child element as the first child of the target.
func (e HtmlElement) PrependChild(child HtmlElement) {
	e.Value.Call("prepend", child.Value)
}

// Inserts sibling element before target element.
func (e HtmlElement) InsertBefore(sibling HtmlElement) {
	e.Value.Call("before", sibling.Value)
}

// Inserts sibling element after target element.
func (e HtmlElement) InsertAfter(sibling HtmlElement) {
	e.Value.Call("after", sibling.Value)
}

// Replaces target element with another.
func (e HtmlElement) ReplaceWith(new HtmlElement) {
	e.Value.Call("replaceWith", new.Value)
}

// TRAVERSAL

// Returns parent element of the target element.
// Returns false if element is not attached to the DOM or if called on the document.
func (e HtmlElement) Parent() (HtmlElement, bool) {
	v := e.Value.Get("parentElement")
	if !v.Truthy() {
		var zero js.Value
		return HtmlElement{zero}, false
	}
	return HtmlElement{v}, true
}

// Returns slice of children of the target element. Returns empty slice if there are not children.
func (e HtmlElement) Children() []HtmlElement {
	children := e.Value.Get("children")
	length := children.Get("length").Int()

	result := make([]HtmlElement, 0, length)

	for i := range length {
		child := children.Index(i)
		result = append(result, HtmlElement{child})
	}
	return result
}

// Returns first child if any.
func (e HtmlElement) FirstChild() (HtmlElement, bool) {
	c := e.Children()
	if len(c) == 0 {
		var zero HtmlElement
		return zero, false
	}

	return c[0], true
}

// Returns last child if any.
func (e HtmlElement) LastChild() (HtmlElement, bool) {
	c := e.Children()
	if len(c) == 0 {
		var zero HtmlElement
		return zero, false
	}

	return c[len(c)-1], true
}

// Returns count of children.
func (e HtmlElement) ChildElementCount() int {
	c := e.Children()
	return len(c)
}

// Returns next ('below') sibling.
func (e HtmlElement) NextElementSibling() (HtmlElement, bool) {
	c := e.Value.Get("nextElementSibling")
	if !c.Truthy() {
		var zero HtmlElement
		return zero, false
	}
	return HtmlElement{c}, true
}

// Returns previous ('above') sibling.
func (e HtmlElement) PreviousElementSibling() (HtmlElement, bool) {
	c := e.Value.Get("previousElementSibling")
	if !c.Truthy() {
		var zero HtmlElement
		return zero, false
	}
	return HtmlElement{c}, true
}

// SHORTCUTS FROM DOM
// Returns element with provided id if any. Scoped to the document.
func (e HtmlElement) GetElementById(id string) (HtmlElement, bool) {
	return GetElementById(e, id)
}

// Returns first element found by CSS selector if any. Scoped to any descendant of the target element.
func (e HtmlElement) QuerySelector(selector string) (HtmlElement, bool) {
	return QuerySelector(e, selector)
}

// Returns all elements found by CSS selector if any. Scoped to any descendant of the target element.
func (e HtmlElement) QuerySelectorAll(selector string) []HtmlElement {
	return QuerySelectorAll(e, selector)
}

// Returns all elements with the provided class if any. Scoped to any descendant of the target element.
func (e HtmlElement) GetElementsByClassName(className string) []HtmlElement {
	return GetElementsByClassName(e, className)
}

// Returns all elements of the provided tag if any. Scoped to any descendant of the target element.
func (e HtmlElement) GetElementsByTagName(tag tag.TagName) []HtmlElement {
	return GetElementsByTagName(e, tag)
}

// Inserts HTML element with specified insertion type relative to the HTML body tag.
func (h HtmlElement) InsertIntoDom(method insert.InsertionMethod) {
	InsertIntoDom(h, method)
}

// Attaches event listener to target element with specified event type.
// Ececutes provided callback. Returns a clean-up function used to remove the event listener.
func (e HtmlElement) AddEventListener(eventType event.EventType, f func()) (cleanup func()) {
	return AddEventListener(e, eventType, f)
}
