//go:build js && wasm

package lib

import (
	"syscall/js"
)

type ReactiveElement[T any] struct {
	element  js.Value
	state    StateGetter[T]
	setState StateSetter[T]
}

type StateGetter[T any] func() T
type StateSetter[T any] func(T)

type Handler[T any] func(js.Value, []js.Value, StateGetter[T], StateSetter[T]) any

func NewReactiveElement[T any](htmlElement js.Value, stateGetter StateGetter[T], stateSetter StateSetter[T]) ReactiveElement[T] {
	return ReactiveElement[T]{
		element:  htmlElement,
		state:    stateGetter,
		setState: stateSetter,
	}
}

func UseState[T any](initial T) (StateGetter[T], StateSetter[T]) {
	state := initial
	return func() T { return state }, func(new T) {
		state = new
	}
}

func AddElement[T any](elementType string, value T) js.Value {
	document := js.Global().Get("document")
	element := document.Call("createElement", elementType)
	element.Set("innerHTML", value)
	document.Get("body").Call("appendChild", element)
	return element
}

// func AddReactiveElement[T any](elementType string, value T) ReactiveElement[T] {
// 	element := AddElement(elementType, value)
// 	state, setState := UseState(value)
// 	return NewReactiveElement(element, state, setState)
// }

func AddEventListener[T any](element ReactiveElement[T], eventType string, handler Handler[T]) {
	element.element.Call("addEventListener", eventType, js.FuncOf(func(this js.Value, args []js.Value) any {
		return handler(this, args, element.state, element.setState)
	}))
}

// type Handler struct {

// }
