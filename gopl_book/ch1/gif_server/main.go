// Server that responds to a client's GET by a lissajous GIF
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

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
	port         = "8000"
)

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: programm.go r | l")
		os.Exit(1)
	}

	var addr string
	env := strings.ToLower(os.Args[1])

	// This can probably be a map instead?
	if env == "r" {
		addr = "34.23.191.23"
	} else if env == "l" {
		addr = "localhost"
	} else {
		fmt.Fprintf(os.Stderr, "Unexpected env: %s\n", env)
		os.Exit(1)
	}

	server := addr + ":" + port

	gifHandler := func(w http.ResponseWriter, r *http.Request) {
		// Parse client params to lissajous input
		clientParams := r.URL.Query()
		gifParams := prepareLissajousParams(clientParams)

		lissajous(w, gifParams)
	}
	http.HandleFunc("/", gifHandler) // each request calls our handler func defined below
	log.Fatal(http.ListenAndServe(server, nil))
}

func lissajous(out io.Writer, prms map[string]int) {
	// Local const: inputs that control GIF visuals
	// Compared to the original lissajous func, some of the params are pulled from Client
	const res = 0.001 // angular resolution (whatever it means)

	freq := rand.Float64() * 3.0 // relative frequency of t oscillator
	// LoopCount controls # of repetitive animations
	anim := gif.GIF{LoopCount: prms["nframes"]}
	phase := 0.0 // phase difference (no idea what it means :()
	// Loop, tbd what it does
	for i := 0; i < prms["nframes"]; i++ {
		size := prms["size"]
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// Another loop, why?
		for t := 0.0; t < float64(prms["cycles"])*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			colIndex := pickRandColor(palette)
			sizeFloat := float64(size)
			img.SetColorIndex(size+int(x*sizeFloat+0.5), size+int(y*sizeFloat+0.5), colIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, prms["delay"])
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

func prepareLissajousParams(input map[string][]string) map[string]int {

	// Our final data container with default vals
	output := map[string]int{
		"cycles":  5,   // # of revolutions
		"size":    100, // image canvas covers
		"nframes": 64,  // # of animation frames
		"delay":   8,   // delay between frames in 10ms units
	}

	// Process input map by pulling data from existing keys with non-zero vals
	for k, v := range input {
		if len(v) == 0 {
			continue
		}
		// Parse string to int and store within output
		i, err := strconv.Atoi(v[0])
		if err != nil {
			continue // Some logging should happen here, but we ignore it :(
		}
		output[k] = i
	}

	return output
}
