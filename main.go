//go:build js && wasm

package main

import "github.com/Robert-Safin/go-wasm/dom"

func main() {

	btn1 := dom.CreateElement(dom.ButtonTag)
	btn2 := dom.CreateElement(dom.ButtonTag)

	btn1.SetProp(dom.InnerHTMLProp, "one")
	btn2.SetProp(dom.InnerHTMLProp, "two")

	btn1.Insert(dom.AppendChildMethod)

	btn1.AddEvent(dom.ClickEvent, func() {
		btn1.ReplaceWith(btn2)
	})

	select {}
}
