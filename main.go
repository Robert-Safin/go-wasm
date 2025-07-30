//go:build js && wasm

package main

import (
	"github.com/Robert-Safin/go-wasm/dom"
	"github.com/Robert-Safin/go-wasm/dom/types/attribute"
	"github.com/Robert-Safin/go-wasm/dom/types/insert"
	"github.com/Robert-Safin/go-wasm/dom/types/tag"
)

func main() {

	btn := dom.CreateElement(tag.Button)
	btn.Insert(insert.AppendChild)
	btn.SetAttribute(attribute.InnerHTML, "click")

	select {}
}
