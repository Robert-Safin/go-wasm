//go:build js && wasm

package examples

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/signal"
)

type Person struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
}

func main() {

	btn := dom.CreateElement(dom.ButtonTag)
	btn.SetAttribute(dom.InnerHTMLAttribute, "next")
	btn.Insert(dom.AppendChildMethod)
	p := dom.CreateElement(dom.PTag)
	p.Insert(dom.AppendChildMethod)

	count := signal.NewSignal(0)
	person := signal.NewSignal(Person{})

	_ = btn.AddEvent(dom.ClickEvent, func() {
		go func() {
			resp, _ := http.Get("https://swapi.info/api/people/" + strconv.Itoa(count.Get()))
			var p Person
			json.NewDecoder(resp.Body).Decode(&p)
			person.Set(p)
		}()
		count.Set(count.Get() + 1)
	})

	signal.Effect(func() {
		p.SetAttribute(dom.InnerHTMLAttribute, person.Get().Name)
	})

	_ = p.AddEvent(dom.ClickEvent, func() {
		btn.Delete()
	})

	select {}
}
