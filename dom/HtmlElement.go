//go:build js && wasm

package dom

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
)

type HtmlElement struct {
	Value js.Value
}

// SHORTCUTS FROM DOM PACKAGE
func (h HtmlElement) InsertIntoDom(method insert.InsertionMethod) bool {
	return InsertIntoDom(h, method)
}
func (e HtmlElement) AddEventListener(eventType event.EventType, f func()) (cleanup func()) {
	return AddEventListener(e, eventType, f)
}

// GETTERS AND SETTER
func (h HtmlElement) GetAttribute(prop attribute.AttributeName) (js.Value, bool) {
	v := h.Value.Get(string(prop))
	if v.IsUndefined() {
		var zero js.Value
		return zero, false
	}
	return v, true
}
func (h HtmlElement) SetAttribute(prop attribute.AttributeName, value string) bool {
	ok := true
	defer func() {
		if r := recover(); r != nil {
			ok = false
			fmt.Println("Recovered from panic during property setting:", r)
		}
	}()
	h.Value.Set(prop.String(), value)
	return ok
}
func (h HtmlElement) SetAttributeMap(props map[attribute.AttributeName]string) bool {
	for k, v := range props {
		ok := h.SetAttribute(k, v)
		if !ok {
			return false
		}
	}
	return true
}
func (e HtmlElement) SetStyles(styles map[string]string) {
	joined := ""
	for k, v := range styles {
		joined += k + ":" + v + ";"
	}
	e.SetAttribute(attribute.Style, joined)
}

// UPDATING
func (e HtmlElement) Delete() {
	e.Value.Call("remove")
}
func (e HtmlElement) ReplaceWith(new HtmlElement) {
	e.Value.Call("replaceWith", new.Value)
}
func (e HtmlElement) AddClasses(classNames ...string) {
	js, _ := e.GetAttribute(attribute.ClassName)
	classes := js.String()

	if classes != "" {
		classes += " "
	}

	for _, c := range classNames {
		classes += c + " "
	}

	if classes != "" {
		classes = classes[:len(classes)-1]

	}
	e.SetAttribute(attribute.ClassName, classes)
}
func (e HtmlElement) RemoveClasses(classNames ...string) {
	toDelete := make(map[string]bool, len(classNames))
	for _, cls := range classNames {
		toDelete[cls] = true
	}

	val, _ := e.GetAttribute(attribute.ClassName)
	existing := strings.Fields(val.String())

	var kept []string
	for _, cls := range existing {
		if !toDelete[cls] {
			kept = append(kept, cls)
		}
	}

	e.SetAttribute(attribute.ClassName, strings.Join(kept, " "))
}

// TRAVERSAL
func (e HtmlElement) Parent() (HtmlElement, bool) {
	v := e.Value.Get("parentElement")
	if v.IsNull() || v.IsUndefined() {
		var zero js.Value
		return HtmlElement{zero}, false
	}
	return HtmlElement{v}, true
}
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
func (e HtmlElement) FirstChild() (HtmlElement, bool) {
	c := e.Children()
	if len(c) == 0 {
		var zero HtmlElement
		return zero, false
	}

	return c[0], true
}
func (e HtmlElement) LastChild() (HtmlElement, bool) {
	c := e.Children()
	if len(c) == 0 {
		var zero HtmlElement
		return zero, false
	}

	return c[len(c)-1], true
}

// nextElementSibling
// previousElementSibling
// childElementCount
func (e HtmlElement) ChildElementCount() int {
	c := e.Children()
	return len(c)
}

// SCOPED QUERIES
// element.querySelector
// element.querySelectorAll
// element.getElementsByClassName
// element.getElementsByTagName
