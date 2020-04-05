package main

import (
	"fmt"
	"math/cmplx"
	"syscall/js"
)

const ITTERATION_LIMIT = 50

var (
	// js.Value can be any JS object/type/constructor
	window, doc, body, canvas, ctx, img, pixels js.Value
	windowSize                                  struct{ w, h float64 }
)

func init() {
	println("Initializing")
	window = js.Global()
	doc = window.Get("document")
	body = doc.Get("body")

	windowSize.h = window.Get("innerHeight").Float()
	windowSize.w = window.Get("innerWidth").Float()

	canvas = doc.Call("createElement", "canvas")
	canvas.Set("height", windowSize.h)
	canvas.Set("width", windowSize.w)
	body.Call("appendChild", canvas)

	ctx = canvas.Call("getContext", "2d")

	println("Finished Initializing")
}

func makeItPink(ctx js.Value) {
	fmt.Println("Making it pink!")
	img = ctx.Call("getImageData", 0, 0, windowSize.w, windowSize.h)
	pixels = img.Get("data")

	for x := 0; x < int(windowSize.w); x++ {
		for y := 0; y < int(windowSize.h); y++ {
			index := (y*int(windowSize.w) + x) * 4
			cx := ((float64(x) - windowSize.w/2) / windowSize.w) * 4
			cy := ((float64(y) - windowSize.h/2) / windowSize.h) * 4

			color := float64(mandelbrot(cx, cy, 2))
			l := float64(ITTERATION_LIMIT)
			green := 255 - ((l-color)/l)*255.

			pixels.SetIndex(index, 0)
			pixels.SetIndex(index+1, green)
			pixels.SetIndex(index+2, 0)
			pixels.SetIndex(index+3, 255)
		}
	}

	img.Set("data", pixels)
	ctx.Call("putImageData", img, 0, 0)
	fmt.Println("It's PINK")
}

func mandelbrot(x, y, threshold float64) int {
	c0 := complex(x, y)
	c := complex(0, 0)
	for i := 0; i < ITTERATION_LIMIT; i++ {
		if cmplx.Abs(c) > threshold {
			return i
		}

		c = c*c + c0
	}

	return 0
}

func test(this js.Value, args []js.Value) interface{} {
	println("testing this function")
	return nil
}

func registerCallbacks() {
	js.Global().Set("test", js.FuncOf(test))
}

func main() {
	c := make(chan struct{}, 0)

	go makeItPink(ctx)
	println("WASM Go Initialized")
	registerCallbacks()

	<-c
}
