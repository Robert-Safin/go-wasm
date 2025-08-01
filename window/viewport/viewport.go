//go:build js && wasm

package viewport

import "syscall/js"

// Viewport width
func InnerWidth() int {
	return js.Global().Get("window").Get("innerWidth").Int()
}

// Viewport height
func InnerHeight() int {
	return js.Global().Get("window").Get("innerHeight").Int()
}
