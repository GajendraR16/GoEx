package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r.URL.Query())
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}

func lissajous(out io.Writer, m url.Values) {

	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := uint8(1 + i%(len(palette)-1))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // red
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // blue
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF}}

const (
	whiteIndex = 0
	blackIndex = 1
)
