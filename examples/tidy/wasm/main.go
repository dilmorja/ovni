// +build js,wasm
package main

import(
	"syscall/js"
)

func main() {
	header := js.Global().Call("createElement", "h1")
	header.Call("innerHTML", "Hello world!")
}