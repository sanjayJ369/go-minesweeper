package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var doc js.Value

func add(this js.Value, i []js.Value) any {
	val1Str := doc.Call("getElementById", i[0].String()).Get("value").String()
	val2Str := doc.Call("getElementById", i[1].String()).Get("value").String()

	val1, _ := strconv.Atoi(val1Str)
	val2, _ := strconv.Atoi(val2Str)

	res := val1 + val2

	doc.Call("getElementById", i[2].String()).Set("textContent", js.ValueOf(res))

	fmt.Println(res)
	return nil
}

func subtract(this js.Value, i []js.Value) any {
	val1Str := doc.Call("getElementById", i[0].String()).Get("value").String()
	val2Str := doc.Call("getElementById", i[1].String()).Get("value").String()

	val1, _ := strconv.Atoi(val1Str)
	val2, _ := strconv.Atoi(val2Str)

	res := val1 - val2

	doc.Call("getElementById", i[2].String()).Set("textContent", js.ValueOf(res))

	fmt.Println(res)
	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)
	doc = js.Global().Get("document")
	registerCallbacks()
	<-c
}
