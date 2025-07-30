//go:build js && wasm

package examples

import (
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

func main1() {

	btn := dom.CreateElement(tag.ButtonTag)
	btn.SetAttribute(attribute.InnerHTMLAttribute, "click me")
	btn.Insert(insert.AppendChildMethod)

	p := dom.CreateElement(tag.PTag)
	p.SetAttribute(attribute.InnerHTMLAttribute, "0")
	p.Insert(insert.AppendChildMethod)

	btn.AddEvent(event.ClickEvent, func() {
		v, _ := p.GetAttribute(attribute.InnerHTMLAttribute)
		n, _ := strconv.Atoi(v.String())
		_ = p.SetAttribute(attribute.InnerHTMLAttribute, strconv.Itoa(n+1))
	})

	// Keep Go alive
	select {}
}
