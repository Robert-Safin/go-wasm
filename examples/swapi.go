//go:build js && wasm

package examples

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

	btn := dom.CreateElement(tag.ButtonTag)
	btn.SetAttribute(attribute.InnerHTMLAttribute, "next")
	btn.Insert(insert.AppendChildMethod)
	p := dom.CreateElement(tag.PTag)
	p.Insert(insert.AppendChildMethod)

	count := signal.NewSignal(0)
	person := signal.NewSignal(Person{})

	_ = btn.AddEvent(event.ClickEvent, func() {
		go func() {
			resp, _ := http.Get("https://swapi.info/api/people/" + strconv.Itoa(count.Get()))
			var p Person
			json.NewDecoder(resp.Body).Decode(&p)
			person.Set(p)
		}()
		count.Set(count.Get() + 1)
	})

	signal.Effect(func() {
		p.SetAttribute(attribute.InnerHTMLAttribute, person.Get().Name)
	})

	_ = p.AddEvent(event.ClickEvent, func() {
		btn.Delete()
	})

	select {}
}
