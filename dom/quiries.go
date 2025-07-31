//go:build js && wasm

package dom

import (
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

// Returns element with provided id if any. Scoped to the document.
func GetElementById(target HtmlElement, id string) (HtmlElement, bool) {
	element := target.Value.Call("getElementById", id)
	if !element.Truthy() {
		var zero HtmlElement
		return zero, false
	}
	return HtmlElement{element}, true
}

// Returns elements with provided class if any. Scoped to the document.
func GetElementsByClassName(target HtmlElement, className string) []HtmlElement {
	elements := target.Value.Call("getElementsByClassName", className)
	res := []HtmlElement{}
	for i := range elements.Length() {
		res = append(res, HtmlElement{elements.Index(i)})
	}
	return res
}

// Returns elements with provided tag if any
func GetElementsByTagName(target HtmlElement, tag tag.TagName) []HtmlElement {
	elements := target.Value.Call("getElementsByTagName", tag.String())
	res := []HtmlElement{}
	for i := range elements.Length() {
		res = append(res, HtmlElement{elements.Index(i)})
	}
	return res
}

// Returns first element found by CSS selector if any. Scoped to any descendant of the target element.
func QuerySelector(target HtmlElement, selector string) (HtmlElement, bool) {
	element := target.Value.Call("querySelector", selector)
	if !element.Truthy() {
		var zero HtmlElement
		return zero, false
	}
	return HtmlElement{element}, true
}

// Returns all elements found by CSS selector if any. Scoped to any descendant of the target element.
func QuerySelectorAll(target HtmlElement, selector string) []HtmlElement {
	elements := target.Value.Call("querySelectorAll", selector)
	res := []HtmlElement{}
	for i := range elements.Length() {
		res = append(res, HtmlElement{elements.Index(i)})
	}
	return res
}
