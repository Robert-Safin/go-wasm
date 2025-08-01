# Go WASM bindings for the DOM and some browser APIs
- Query, traverse, insert, update and delete DOM elements.
- Instinctive signal based reactivity system (40 lines).
- Resolver for Promises.
- Full support for fetch API, as well as some other minor ones.

# Why?
- Simple, declarative and idiomatic Go.
- Add client-side interactivity without touching HTML, CSS or JavaScript, all in Go.
- Pre-compile wasm binaries, and serve them from any existing HTML serving backend.

# Primitives
```go
	div := dom.CreateElement(tag.Div)
	dom.InsertIntoDom(div, insert.AppendChild)

	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "click me")
	div.AppendChild(btn)

	p := dom.CreateElement(tag.P)
	p.SetStyles(map[string]string{
		"font-size": "40px",
		"color":     "red",
	})
	div.InsertAfter(p)

	count, _ := signal.NewSignal(0, nil)

	p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))

	count.Effect(func() {
		p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))
	})

	btn.AddEventListener(event.Click, func() {
		count.Set(count.Get() + 1)
	})
```

# Another example in detail
Compiled binary 3.2mb
```go
// required build tag.
//go:build js && wasm

package main

import (
	"encoding/json"
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

type Person struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Mass   string `json:"mass"`
}

func main() {
	// find container
	doc := dom.GetDocument()
	container, ok := dom.GetElementById(doc, "container")

	// or create if it doesnt exist
	if !ok {
		container = dom.CreateElement(tag.Div)
		container.SetAttribute(attribute.ID, "container")
		dom.InsertIntoDom(container, insert.After)
	}

	// set some styles into style attribute
	container.SetStyles(map[string]string{
		"background-color": "grey",
		"width":            "50%",
		"height":           "100%",
	})
	// ship with .css files or fetch tailwind stylesheet
	container.AddClasses("flex", "flex-row")

	// Siagnals as reactivity system, stores generic type.
	// Requires initial value and equality funciton.
	// Equality func used to prevent re-renders when new state == prev state.
	// Equality func can be nil, in which case all updates trigger re-renders.
	// Return clean up function, that empties underlying map of effect on a signal.
	count, killSig := signal.NewSignal[int](0, func(a int, b int) bool { return a == b })
	defer killSig()

	// element to render signal value
	p := dom.CreateElement(tag.P)
	p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))
	container.AppendChild(p)

	// this element will store value of some decoded JSON.
	p2 := dom.CreateElement(tag.P)
	container.AppendChild(p2)

	// Add an effect to the signal. Effect is a function that will run if
	// the value of signal changes (count.Set() in this case) unless prevented by Equality func.
	// Any number of effects can be registered to a signal.
	// .Effect returns cleanup function, that will delete only that effect from a signal's map.
	unsub := count.Effect(func() {
		// update 'counter' on signal change
		p.SetAttribute(attribute.InnerHTML, strconv.Itoa(count.Get()))
		// fire a http request and decode it into HTML content
		fetch.Fetch("GET", "https://swapi.info/api/people/"+strconv.Itoa(count.Get()), map[string]string{}, "",
			// promise resolved
			func(bytes []byte) {
				var person Person
				json.Unmarshal(bytes, &person)
				p2.SetAttribute(attribute.InnerHTML, person.Name)

			},
			// promise resolved rejected
			func(e error.Error) {
				console.Log(e.Value)
			})
	})
	defer unsub()

	// add button that will trigger signal updatess
	btn := dom.CreateElement(tag.Button)
	btn.SetAttribute(attribute.InnerHTML, "click me")
	container.PrependChild(btn)

	// attach event listener.
	// returns a clean-up function to remove listener from DOM.
	cleanup := btn.AddEventListener(event.Click, func() {
		count.Set(count.Get() + 1)
	})
	defer cleanup()
	// keep go runtime alive
	select {}
}
```

# Set-up
##### Compile
`GOOS=js GOARCH=wasm go build -o main.wasm ./main.go`
##### 'Glue'
Load compiled wasm binary inside skeleton html as seen in `index.html` and also include `wasm_exec.js`.
##### Spin up static fileserver
`serve .`
