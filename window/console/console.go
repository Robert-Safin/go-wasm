//go:build js && wasm

package console

import (
	"syscall/js"
)

// supports: nil, bool, integer, float, string, []any, map[string]any, js.Value, js.Func
// otherwise causes panic
func Log(args ...any) {
	js.Global().Get("console").Call("log", args...)

}
