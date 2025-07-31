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
	div := dom.CreateElement(tag.Div)
	dom.InsertIntoDom(div, insert.AppendChild)

	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "click me")
	div.AppendChild(btn)

	p := dom.CreateElement(tag.P)
	p.SetAttributeMap(map[attribute.AttributeName]string{
		attribute.InnerHTML: "0",
		attribute.ID:        "count",
	})
	p.SetStyles(map[string]string{
		"font-size": "40px",
		"color":     "red",
	})

	div.InsertAfter(p)

	btn.AddEventListener(event.Click, func() {
		s, _ := p.GetAttribute(attribute.InnerHTML)
		n, _ := strconv.Atoi(s)
		p.SetAttribute(attribute.InnerHTML, strconv.Itoa(n+1))
	})

	select {}
}
