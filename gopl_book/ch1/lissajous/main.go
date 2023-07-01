package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// Declaring global variables
// var palette = []color.Color{color.White, color.Black} // This is the default pallete
// 1.5 Change the color scheme to green on black - DONE
var (
	bgCol           = color.RGBA{13, 2, 8, 10}
	curveCol        = color.RGBA{0, 143, 17, 10}
	extraCurveCol   = color.RGBA{168, 88, 88, 10}
	oneMoreCurveCol = color.RGBA{9, 65, 152, 10}
	palette         = []color.Color{bgCol, curveCol, extraCurveCol, oneMoreCurveCol}
)

const (
	bgIndex      = 0
	mainColIndex = 1
	extraCol1    = 2
	extraCol2    = 3
)

func main() {
	// Call our custom func to generate GIF to STDOUT
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	// Local const: inputs that control GIF visuals
	const (
		cycles  = 5     // # of revolutions
		res     = 0.001 // angular resolution (whatever it means)
		size    = 100   // image canvas covers
		nframes = 64    // # of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of t oscillator
	// LoopCount controls # of repetitive animations
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference (no idea what it means :()
	// Loop, tbd what it does
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// Another loop, why?
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 1.6. Modify the code to produce image in multiple colors - DONE!
			colIndex := pickRandColor(palette)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colIndex) // This creates curves of the same color all the time
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // Intentionally ignore errors!

}

func pickRandColor(palette []color.Color) uint8 {
	// Random seed for choosing our color
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	// Generation
	res := r.Intn(len(palette[1:])) + 1
	return uint8(res)
}
