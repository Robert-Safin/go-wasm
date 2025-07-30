//go:build js && wasm

package examples

import (
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
)

func main1() {

	btn := dom.CreateElement(dom.ButtonTag)
	btn.SetAttribute(dom.InnerHTMLAttribute, "click me")
	btn.Insert(dom.AppendChildMethod)

	p := dom.CreateElement(dom.PTag)
	p.SetAttribute(dom.InnerHTMLAttribute, "0")
	p.Insert(dom.AppendChildMethod)

	btn.AddEvent(dom.ClickEvent, func() {
		v, _ := p.GetAttribute(dom.InnerHTMLAttribute)
		n, _ := strconv.Atoi(v.String())
		_ = p.SetAttribute(dom.InnerHTMLAttribute, strconv.Itoa(n+1))
	})

	// Keep Go alive
	select {}
}
