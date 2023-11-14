// Server responds to http requests with varying lissajous GIFs
// This version allows color to be specified as a URL query param
// For example, http://localhost:8000/?color=yellow
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0xFF, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
	color.White,
	color.RGBA{255, 252, 127, 0xFF},
}

var colors = map[string]int{
	"green":  1,
	"red":    2,
	"blue":   3,
	"white":  4,
	"yellow": 5,
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	colorParam := r.URL.Query().Get("color")
	fmt.Println(colorParam)
	colorIndex := colors[colorParam]
	if colorIndex == 0 {
		colorIndex = rand.Intn(5) + 1
	}
	lissajous(w, colorIndex)
}

func lissajous(out http.ResponseWriter, colorIndex int) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 1; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
