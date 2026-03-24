package main

import (
	"fmt"
	"strconv"

	"golang.org/x/sys/js"
)

func main() {
	fmt.Println("Go WASM carregado!")

	document := js.Global().Get("document")
	body := document.Get("body")

	container := document.CreateElement("div")
	container.Set("style", "font-family: 'Segoe UI', sans-serif; background: linear-gradient(135deg, #1a1a2e 0%, #16213e 100%); min-height: 100vh; display: flex; flex-direction: column; justify-content: center; align-items: center; color: #fff;")

	title := document.CreateElement("h1")
	title.Set("textContent", "Olá do Go WASM!")
	title.Set("style", "color: #00d4ff; font-size: 3rem; margin-bottom: 1rem;")

	subtitle := document.CreateElement("p")
	subtitle.Set("textContent", "Esta página foi renderizada completamente pelo Go compilado para WebAssembly!")
	subtitle.Set("style", "font-size: 1.3rem; margin-bottom: 2rem; max-width: 600px; text-align: center;")

	button := document.CreateElement("button")
	button.Set("textContent", "Clique em mim!")
	button.Set("style", "background: #00d4ff; color: #1a1a2e; border: none; padding: 1rem 2rem; font-size: 1.2rem; border-radius: 30px; cursor: pointer; font-weight: bold; transition: transform 0.3s;")

	counter := document.CreateElement("div")
	counter.Set("textContent", "Cliques: 0")
	counter.Set("style", "margin-top: 1.5rem; font-size: 1.5rem; color: #00d4ff;")

	clickCount := 0

	callback := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		clickCount++
		counter.Set("textContent", "Cliques: "+strconv.Itoa(clickCount))
		button.Get("style").Set("transform", "scale(1.1)")
		return nil
	})

	button.Call("addEventListener", "click", callback)

	container.AppendChild(title)
	container.AppendChild(subtitle)
	container.AppendChild(button)
	container.AppendChild(counter)
	body.AppendChild(container)

	<-make(chan struct{})
}
