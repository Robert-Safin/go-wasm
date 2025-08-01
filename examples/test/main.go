//go:build js && wasm

package main

import (
	"fmt"
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
	"github.com/Robert-Safin/go-wasm/error"
	"github.com/Robert-Safin/go-wasm/signal"
	"github.com/Robert-Safin/go-wasm/window/console"
	"github.com/Robert-Safin/go-wasm/window/fetch"
)

func main() {
	div := dom.CreateElement(tag.Div)
	dom.InsertIntoDom(div, insert.AppendChild)

	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "click me")
	div.AppendChild(btn)

	p := dom.CreateElement(tag.P)
	responseP := dom.CreateElement(tag.P)
	div.PrependChild(responseP)

	count, _ := signal.NewSignal(0, func(a int, b int) bool {
		return a == b
	})

	p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))

	count.Effect(func() {
		p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))
		fetch.Fetch("get", "http://127.0.0.1:3001/", map[string]string{}, "", func(b []byte) {
			console.Log(string(b))
		}, func(e error.Error) {
			fmt.Println(e)
			fmt.Println(e.Message())
			console.Log(e.Message())
		})

	})

	div.InsertAfter(p)

	btn.AddEventListener(event.Click, func() {
		count.Set(count.Get() + 1)
	})

	select {}
}
