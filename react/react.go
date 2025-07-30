//go:build js && wasm

package react

import "syscall/js"

type ReactiveElement[T any] struct {
	Element  js.Value
	State    StateGetter[T]
	SetState StateSetter[T]
}

type StateGetter[T any] func() T
type StateSetter[T any] func(T)

func NewReactiveElement[T any](htmlElement js.Value, stateGetter StateGetter[T], stateSetter StateSetter[T]) ReactiveElement[T] {
	return ReactiveElement[T]{
		Element:  htmlElement,
		State:    stateGetter,
		SetState: stateSetter,
	}
}

func UseState[T any](initial T) (StateGetter[T], StateSetter[T]) {
	state := initial
	return func() T { return state }, func(new T) {
		state = new
	}
}
