//go:build js && wasm

package dom

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

func CreateElement(elementType tag.TagName) HtmlElement {
	document := js.Global().Get("document")
	element := document.Call("createElement", elementType.String())
	return HtmlElement{element}
}
