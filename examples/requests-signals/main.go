//go:build js && wasm

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
	"github.com/Robert-Safin/go-wasm/signal"
)

type Person struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
}

func main() {

	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "next")
	btn.InsertIntoDom(insert.AppendChild)
	p := dom.CreateElement(tag.P)
	p.SetAttribute(attribute.InnerHTML, "Nothing")
	p.InsertIntoDom(insert.AppendChild)

	count := signal.NewSignal(0)
	person := signal.NewSignal(Person{})

	_ = btn.AddEventListener(event.Click, func() {
		go func() {
			resp, _ := http.Get("https://swapi.info/api/people/" + strconv.Itoa(count.Get()))
			var p Person
			json.NewDecoder(resp.Body).Decode(&p)
			person.Set(p)
		}()
		count.Set(count.Get() + 1)
	})

	signal.Effect(func() {
		p.SetAttribute(attribute.InnerHTML, person.Get().Name)
	})

	_ = p.AddEventListener(event.Click, func() {
		btn.Delete()
	})

	select {}
}
