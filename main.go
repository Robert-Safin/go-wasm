//go:build js && wasm

package main

import (
	"fmt"

	"github.com/Robert-Safin/go-wasm/dom"
)

func main() {
	// state, setState := react.UseState(0)
	// button := dom.AddElement(typed.ButtonTag.String(), state())
	// buttonReactive := react.NewReactiveElement(button, state, setState)

	// dom.AddEventListener(buttonReactive, typed.ClickEvent, func(target js.Value, args []js.Value, getState func() int, setState func(int)) any {
	// 	return nil
	// })
	el, ok := dom.GetElementById("hello")
	fmt.Println(el, ok)

	// Keep Go alive
	select {}
}
