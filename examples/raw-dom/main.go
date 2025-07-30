//go:build js && wasm

package main

import (
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

func main() {

	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "click me")
	btn.InsertIntoDom(insert.AppendChild)

	p := dom.CreateElement(tag.P)
	p.SetAttribute(attribute.InnerHTML, "0")
	p.InsertIntoDom(insert.AppendChild)

	btn.AddEventListener(event.Click, func() {
		v, _ := p.GetAttribute(attribute.InnerHTML)
		n, _ := strconv.Atoi(v.String())
		_ = p.SetAttribute(attribute.InnerHTML, strconv.Itoa(n+1))
	})

	// Keep Go alive
	select {}
}
