//go:build js && wasm

package main

import (
	"github.com/Robert-Safin/go-wasm/dom"
)

func main() {

	el := dom.CreateElement(dom.H1Tag)
	el.SetProp(dom.InnerHTMLProp, "hello")
	//fmt.Println(el.GetProp(typed.InnerHTMLProp))

	el.Insert(dom.AppendChildMethod)

	// Keep Go alive
	select {}
}
