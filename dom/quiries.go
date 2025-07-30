//go:build js && wasm

package dom

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

func GetElementById(id string) (HtmlElement, bool) {
	document := js.Global().Get("document")
	element := document.Call("getElementById", id)

	if element.IsUndefined() || element.IsNull() {
		var zero HtmlElement
		return zero, false
	}
	return HtmlElement{element}, true
}

func GetElementsByClassName(class string) ([]HtmlElement, bool) {
	document := js.Global().Get("document")
	elements := document.Call("getElementsByClassName", class)

	length := elements.Get("length").Int()

	if length == 0 {
		return nil, false
	}

	result := make([]HtmlElement, 0, length)

	for i := range length {
		element := elements.Index(i)
		result = append(result, HtmlElement{element})
	}

	return result, true
}

func GetElementsByTagName(tag tag.TagName) ([]HtmlElement, bool) {
	document := js.Global().Get("document")
	elements := document.Call("getElementsByTagName", tag.String())

	length := elements.Get("length").Int()

	if length == 0 {
		return nil, false
	}

	result := make([]HtmlElement, 0, length)

	for i := range length {
		element := elements.Index(i)
		result = append(result, HtmlElement{element})
	}

	return result, true
}

// TODO
// querySelector(selector)
// querySelectorAll(selector)
// SCOPED QUERIES
// element.querySelector
// element.querySelectorAll
// element.getElementsByClassName
// element.getElementsByTagName
// NODE TRAVERSAL
// parentNode, firstChild, lastChild, nextSibling, previousSibling, children
