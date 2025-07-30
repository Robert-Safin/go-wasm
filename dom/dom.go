//go:build js && wasm

package dom

import (
	"syscall/js"
)

func CreateElement(elementType TagName) HtmlElement {
	document := js.Global().Get("document")
	element := document.Call("createElement", elementType.String())
	return HtmlElement{element}
}
