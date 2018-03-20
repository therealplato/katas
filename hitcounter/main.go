package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("vim-go")
	src, err := os.Open("numbers.png")
	if err != nil {
		log.Fatal(err)
	}

	in, err := png.Decode(src)
	if err != nil {
		log.Fatal(err)
	}

	e := &endpoint{
		counts: make(map[string]int),
		source: in,
	}

	server := &http.Server{
		Addr:    "0.0.0.0:7777",
		Handler: e,
	}
	server.ListenAndServe()

}

func num(in image.Image, digit int) image.Image {
	var (
		sr image.Rectangle
		dr image.Rectangle
	)
	switch digit {
	case 0:
		sr = ZERO
	case 1:
		sr = ONE
	case 2:
		sr = TWO
	case 3:
		sr = THREE
	case 4:
		sr = FOUR
	case 5:
		sr = FIVE
	case 6:
		sr = SIX
	case 7:
		sr = SEVEN
	case 8:
		sr = EIGHT
	case 9:
		sr = NINE
	}
	dr = sr.Sub(sr.Min)
	spew.Dump(sr)
	spew.Dump(dr)
	dst := image.NewGray(dr) // 0,0 .. 100,100
	draw.Draw(dst, dr, in, sr.Min, draw.Src)
	return dst
}

func stack(ins ...image.Image) image.Image {
	h := 100 * len(ins)
	w := 100
	dst := image.NewGray(image.Rect(0, 0, w, h)) // 0,0 .. 100,100

	for i, digit := range ins {
		sr := digit.Bounds()
		dr := image.Rect(0, i*100, 100, i*100+100)
		draw.Draw(dst, dr, digit, sr.Min, draw.Src)
	}
	return dst
}
func row(ins ...image.Image) image.Image {
	h := 100
	w := 100 * len(ins)
	dst := image.NewGray(image.Rect(0, 0, w, h)) // 0,0 .. 100,100

	for i, digit := range ins {
		sr := digit.Bounds()
		dr := image.Rect(i*100, 0, i*100+100, 100)
		draw.Draw(dst, dr, digit, sr.Min, draw.Src)
	}
	return dst
}

func counter(in image.Image, n int) image.Image {
	digits := make([]image.Image, 0)
	s := strconv.Itoa(n)
	for _, c := range s {
		d, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		digit := num(in, d)
		digits = append(digits, digit)
	}
	if n%2 == 0 {
		return row(digits...)
	}
	return stack(digits...)
}
