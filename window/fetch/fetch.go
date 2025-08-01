//go:build js && wasm

package fetch

import (
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/error"
	"github.com/Robert-Safin/go-wasm/promise"
)

func Fetch(
	method string,
	url string,
	headers map[string]string,
	body string,
	onSuccess func([]byte),
	onError func(error.Error),
) {
	jsHeaders := js.Global().Get("Object").New()
	for key, value := range headers {
		jsHeaders.Set(key, value)
	}

	opts := js.Global().Get("Object").New()
	opts.Set("method", method)
	opts.Set("headers", jsHeaders)
	if method != "GET" && body != "" {
		opts.Set("body", body)
	}

	p := promise.IntoPromise(js.Global().Call("fetch", url, opts))

	p.Resolve(
		func(resp js.Value) {
			arrayBufferPromise := promise.IntoPromise(resp.Call("arrayBuffer"))
			arrayBufferPromise.Resolve(
				func(buffer js.Value) {
					uint8Array := js.Global().Get("Uint8Array").New(buffer)
					length := uint8Array.Get("length").Int()
					data := make([]byte, length)
					js.CopyBytesToGo(data, uint8Array)
					onSuccess(data)
				},
				func(errMsg error.Error) {
					onError(errMsg)
				},
			)
		},
		func(errMsg error.Error) {
			onError(errMsg)
		},
	)
}
