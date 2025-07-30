//go:build js && wasm

package main

import (
	"fmt"

	"github.com/Robert-Safin/go-wasm/dom"
)

func main() {
	d, _ := dom.GetElementById("outer")
	n := d.ChildElementCount()
	fmt.Println(n)

	// p, ok := btn.Parent()
	// fmt.Println(p, ok)

	select {}
}
