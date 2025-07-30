//go:build js && wasm

package dom

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

func GetDocument() js.Value {
	document := js.Global().Get("document")
	return document
}

func CreateElement(elementType tag.TagName) HtmlElement {
	document := GetDocument()
	element := document.Call("createElement", elementType.String())
	return HtmlElement{element}
}

func AddEventListener(target HtmlElement, eventType event.EventType, f func()) (cleanup func()) {
	handler := js.FuncOf(func(this js.Value, args []js.Value) any {
		f()
		return nil
	})
	target.Value.Call("addEventListener", eventType.String(), handler)
	return handler.Release
}

func InsertIntoDom(element HtmlElement, method insert.InsertionMethod) bool {
	document := js.Global().Get("document")
	document.Get("body").Call(method.String(), element.Value)
	return true
}
