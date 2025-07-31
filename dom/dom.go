//go:build js && wasm

package dom

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

// Retuns the document
func GetDocument() HtmlElement {
	document := js.Global().Get("document")
	return HtmlElement{document}
}

// Creates a blank HTML element of specified tag
func CreateElement(elementType tag.TagName) HtmlElement {
	document := GetDocument()
	element := document.Value.Call("createElement", elementType.String())
	return HtmlElement{element}
}

// Attaches event listener to target element with specified event type.
// Ececutes provided callback. Returns a clean-up function used to remove the event listener.
func AddEventListener(target HtmlElement, eventType event.EventType, f func()) (cleanup func()) {
	handler := js.FuncOf(func(this js.Value, args []js.Value) any {
		f()
		return nil
	})
	target.Value.Call("addEventListener", eventType.String(), handler)
	return handler.Release
}

// Inserts target element relative to the body tag.
func InsertIntoDom(element HtmlElement, method insert.InsertionMethod) {
	document := js.Global().Get("document")
	document.Get("body").Call(method.String(), element.Value)
}
