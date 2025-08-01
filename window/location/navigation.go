//go:build js && wasm

package location

import "syscall/js"

func Navigate(url string) {
	js.Global().Get("window").Get("location").Set("href", url)
}
