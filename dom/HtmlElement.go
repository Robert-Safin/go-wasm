//go:build js && wasm

package dom

import (
	"fmt"
	"syscall/js"
)

type HtmlElement struct {
	Value js.Value
}

func (h HtmlElement) GetAttribute(prop AttributeName) (js.Value, bool) {
	v := h.Value.Get(string(prop))
	if v.IsUndefined() {
		var zero js.Value
		return zero, false
	}
	return v, true
}

func (h HtmlElement) SetAttribute(prop AttributeName, value string) bool {
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

func (h HtmlElement) SetAttributeMap(props map[AttributeName]string) bool {
	for k, v := range props {
		ok := h.SetAttribute(k, v)
		if !ok {
			return false
		}
	}
	return true
}

func (h HtmlElement) Insert(method InsertionMethod) bool {
	document := js.Global().Get("document")
	document.Get("body").Call(method.String(), h.Value)
	return true
}

func (e HtmlElement) AddEvent(eventType EventType, f func()) (cleanup func()) {
	handler := js.FuncOf(func(this js.Value, args []js.Value) any {
		f()
		return nil
	})
	e.Value.Call("addEventListener", eventType.String(), handler)
	return handler.Release
}

func (e HtmlElement) Delete() {
	e.Value.Call("remove")
}

func (e HtmlElement) SetStyles(styles map[string]string) {
	joined := ""
	for k, v := range styles {
		joined += k + ":" + v + ";"
	}
	e.SetAttribute(StyleAttribute, joined)
}

func (e HtmlElement) ReplaceWith(new HtmlElement) {
	e.Value.Call("replaceWith", new.Value)
}

func (e HtmlElement) Children(new HtmlElement) {

}
