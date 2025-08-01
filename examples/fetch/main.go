//go:build js && wasm

package main

import (
	"encoding/json"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/event"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
	"github.com/Robert-Safin/go-wasm/error"
	"github.com/Robert-Safin/go-wasm/window/console"
	"github.com/Robert-Safin/go-wasm/window/fetch"
)

type Person struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
}

func main() {
	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "fetch")
	btn.InsertIntoDom(insert.AppendChild)

	btn.AddEventListener(event.Click, func() {
		fetch.Fetch("GET", "https://swapi.info/api/people", map[string]string{}, "",
			func(res []byte) {
				var people []Person
				json.Unmarshal(res, &people)
				for _, p := range people {
					t := dom.CreateElement(tag.P)
					t.SetAttribute(attribute.InnerHTML, p.Name+" "+p.Height+"cm "+p.Mass+"kg")
					t.InsertIntoDom(insert.AppendChild)
				}

			},
			func(e error.Error) {
				console.Log(e.Value)
			})
	})

	select {}
}
