//go:build js && wasm

package promise

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/error"
)

// Wrapper for implementation
type Promise struct {
	Value js.Value
}

func IntoPromise(v js.Value) *Promise {
	return &Promise{v}
}

func (p *Promise) Resolve(onSuccess func(js.Value), onError func(error.Error)) {
	var thenHandler js.Func
	var catchHandler js.Func

	thenHandler = js.FuncOf(func(this js.Value, args []js.Value) any {
		ok := args[0]
		onSuccess(ok)
		thenHandler.Release()
		catchHandler.Release()
		return nil
	})

	catchHandler = js.FuncOf(func(this js.Value, args []js.Value) any {
		err := args[0]
		onError(error.Error{Value: err})
		thenHandler.Release()
		catchHandler.Release()
		return nil
	})

	p.Value.Call("then", thenHandler).Call("catch", catchHandler)
}
