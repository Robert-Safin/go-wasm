//go:build js && wasm

package examples

import (
	"strconv"
	"syscall/js"

	"github.com/Robert-Safin/go-wasm/signal"
)

func main2() {
	doc := js.Global().Get("document")

	// Create signal
	count := signal.NewSignal(0)

	// Create button
	btn := doc.Call("createElement", "button")
	btn.Set("textContent", "Click me")
	btn.Call("addEventListener", "click", js.FuncOf(func(this js.Value, args []js.Value) any {
		count.Set(count.Get() + 1)
		return nil
	}))
	doc.Get("body").Call("appendChild", btn)

	// Create paragraph bound to signal
	p := doc.Call("createElement", "p")
	signal.Effect(func() {
		p.Set("textContent", "Count: "+strconv.Itoa(count.Get()))
	})
	doc.Get("body").Call("appendChild", p)

	select {} // Keep Go runtime alive
}
