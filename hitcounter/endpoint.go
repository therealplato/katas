package main

import (
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
)

type endpoint struct {
	// store in a git repo?
	// might help to survive cloudburst
	counts map[string]int
	source image.Image
}

// http://localhost:8080/counter/${{identifier}}

func (e *endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.counts[r.URL.Path]++

	out := counter(e.source, e.counts[r.URL.Path])
	err := png.Encode(w, out)
	if err != nil {
		f, _ := os.Create("brokencounter.png")
		png.Encode(f, out) // yeah, right, this is gonna work...
		log.Fatal("what have i done")
	}
	return
}
