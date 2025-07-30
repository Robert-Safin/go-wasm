//go:build js && wasm

package main

import (
	"github.com/Robert-Safin/go-wasm/dom"
)

func main() {

	el, _ := dom.GetElementById("hi")
	el.AddClasses("bob", "alice")
	el.RemoveClasses("bob")

	select {}
}
