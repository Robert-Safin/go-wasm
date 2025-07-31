//go:build js && wasm

package main

import (
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
	"github.com/Robert-Safin/go-wasm/signal"
)

func main() {
	div := dom.CreateElement(tag.Div)
	dom.InsertIntoDom(div, insert.AppendChild)

	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "click me")
	div.AppendChild(btn)

	p := dom.CreateElement(tag.P)

	count, cleanup := signal.NewSignal(0, func(a int, b int) bool {
		return a == b
	})

	p.SetAttributeMap(map[attribute.AttributeName]string{
		attribute.ID:        "count",
		attribute.InnerHTML: strconv.Itoa(count.Get()),
	})

	_ = count.Effect(func() {
		p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))
	})

	p.SetStyles(map[string]string{
		"font-size": "40px",
		"color":     "red",
	})

	div.InsertAfter(p)

	btn.AddEventListener(event.Click, func() {
		count.Set(count.Get() + 1)
		if count.Get() == 10 {
			cleanup()
		}
	})

	select {}
}
