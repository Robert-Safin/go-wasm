//go:build js && wasm

package dom

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/react"
	"github.com/Robert-Safin/go-wasm/typed"
)

func GetElementById(id string) (js.Value, bool) {
	document := js.Global().Get("document")
	element := document.Call("getElementById", id)
	if element.IsNull() {
		var zero js.Value
		return zero, false
	}
	return element, true
}

func AddElement[T any](elementType string, value T) js.Value {
	document := js.Global().Get("document")
	element := document.Call("createElement", elementType)
	element.Set("innerHTML", value)
	document.Get("body").Call("appendChild", element)
	return element
}

type Handler[T any] func(js.Value, []js.Value, react.StateGetter[T], react.StateSetter[T]) any

func AddEventListener[T any](target react.ReactiveElement[T], eventType typed.EventType, handler Handler[T]) {
	target.Element.Call("addEventListener", eventType, func(this js.Value, args []js.Value) any {
		return handler(this, args, target.State, target.SetState)
	})
}
