//go:build js && wasm

package viewport

import "syscall/js"

type Viewport struct {
	Value js.Value
}

func Init() *Viewport {
	v := js.Global().Get("window")
	return &Viewport{v}
}
func (v *Viewport) InnerWidth() int {
	return v.Value.Get("innerWidth").Int()

}
func (v *Viewport) InnerHeight() int {
	return v.Value.Get("innerHeight").Int()
}
