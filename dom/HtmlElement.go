//go:build js && wasm

package dom

import (
	"fmt"
	"syscall/js"
)

type HtmlElement struct {
	js.Value
}

func (h HtmlElement) GetProp(prop PropertyName) (js.Value, bool) {
	v := h.Get(string(prop))
	if v.IsUndefined() {
		var zero js.Value
		return zero, false
	}
	return v, true
}

func (h HtmlElement) SetProp(prop PropertyName, value any) bool {
	ok := true
	defer func() {
		if r := recover(); r != nil {
			ok = false
			fmt.Println("Recovered from panic during property setting:", r)
		}
	}()
	h.Set(prop.String(), value)
	return ok
}

func (h HtmlElement) Insert(method InsertionMethod) bool {
	document := js.Global().Get("document")
	document.Get("body").Call(method.String(), h.Value)
	return true
}
