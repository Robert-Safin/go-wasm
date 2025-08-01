//go:build js && wasm

package console

import (
	"syscall/js"
)

// Supports: nil, bool, integer, float, string, []any, map[string]any, js.Value, js.Func
// Otherwise causes panic.
// When calling with wrapper types e.g. error.Error, dom.HtmlElement, promise.Promise will panic,
// unless unwrapped by calling .Value on the wrapper first, resulting in native JS value being printed.
func Log(args ...any) {
	js.Global().Get("console").Call("log", args...)

}
