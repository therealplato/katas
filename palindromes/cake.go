package main

import (
	"fmt"
	"log"
)

type layer struct {
	i     int // left index
	j     int // right index
	chars []byte
	left  byte
	right byte
	up    *layer
	down  *layer
}

func (l layer) String() string {
	bb := make([]byte, len(l.chars))
	spaceFill(bb)
	bb[l.i] = l.chars[l.i]
	bb[l.j] = l.chars[l.j]
	return string(bb)
}

func build(chars []byte) (bottom *layer) {
	log.Println("building " + string(chars))
	// we'll destructively modify slice header:
	c := make([]byte, len(chars), len(chars))
	copy(c, chars)

	// layers[0] = bottom
	layers := make([]*layer, 0)

	i := 0
	j := len(c)
	for i <= j {
		c = chars[i:j]
		j--
		fmt.Printf("iterating layer, input: %s\n", string(c))
		lay := &layer{
			i:     i,
			j:     j,
			chars: chars,
		}
		i++
		if len(layers) > 0 {
			top := layers[len(layers)-1]
			lay.down = top
			top.up = lay
		}
		layers = append(layers, lay)
	}
	if len(layers) == 0 {
		return nil
	}
	return layers[0]
}

func spaceFill(bb []byte) {
	for i := range bb {
		bb[i] = byte(' ')
	}
}
