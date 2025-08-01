//go:build js && wasm

package clipboard

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/error"
	"github.com/Robert-Safin/go-wasm/promise"
)

// Copies value to clipboard
// Use callabacks to detect failure/success
func Copy(value string, onSuccess func(js.Value), onError func(error.Error)) {
	p := promise.IntoPromise(js.Global().Get("navigator").Get("clipboard").Call("writeText", js.ValueOf(value)))
	p.Resolve(
		func(v js.Value) {
			if onSuccess != nil {
				onSuccess(v)
			}
		},
		func(errMsg error.Error) {
			if onError != nil {
				onError(errMsg)
			}
		},
	)
}
