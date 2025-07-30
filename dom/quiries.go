//go:build js && wasm

package dom

import "syscall/js"

func GetElementById(id string) (js.Value, bool) {
	document := js.Global().Get("document")
	element := document.Call("getElementById", id)
	if element.IsNull() {
		var zero js.Value
		return zero, false
	}
	return element, true
}
