//go:build js && wasm

package dom

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/react"
)

func CreateElement(elementType TagName) HtmlElement {
	document := js.Global().Get("document")
	element := document.Call("createElement", elementType.String())
	return HtmlElement{element}
}

type Handler[T any] func(js.Value, []js.Value, react.StateGetter[T], react.StateSetter[T]) any

func AddEventListener[T any](target react.ReactiveElement[T], eventType EventType, handler Handler[T]) {
	target.Element.Call("addEventListener", eventType, func(this js.Value, args []js.Value) any {
		return handler(this, args, target.State, target.SetState)
	})
}
