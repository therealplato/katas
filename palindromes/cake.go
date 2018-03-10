package main

import (
	"fmt"
	"log"
)

type cake []*layer

func (c *cake) String() string {
	s := ""
	for i := len(*c) - 1; i >= 0; i-- {
		// fmt.Println(c)
		fmt.Println((*c)[i])
		s = s + (*c)[i].String() + "\n"
	}
	return s
}

func (c *cake) bottom() *layer {
	if len(*c) == 0 {
		return nil
	}
	return (*c)[0]
}

type layer struct {
	i     int // left index
	j     int // right index
	chars []byte
	left  byte
	right byte
	up    *layer
	down  *layer
}

func (l *layer) push(top *layer) {
	if l != nil {
		l.up = top
		top.down = l
	}
}

func (l layer) String() string {
	bb := make([]byte, len(l.chars))
	spaceFill(bb)
	bb[l.i] = l.chars[l.i]
	bb[l.j] = l.chars[l.j]
	return string(bb)
}

func bake(chars []byte) cake {
	if len(chars) == 0 {
		return nil
	}
	log.Println("building " + string(chars))
	// we'll destructively modify slice header:
	c := make([]byte, len(chars), len(chars))
	copy(c, chars)

	i := 0
	j := len(c) - 1
	layers := make([]*layer, 0)

	for i <= j {
		layers, i, j = updateLayers(layers, c, i, j)
		fmt.Printf("layers: %v\n", layers)
	}
	_, _ = i, j
	if len(layers) == 0 {
		return nil
	}
	return cake(layers)
}

func spaceFill(bb []byte) {
	for i := range bb {
		bb[i] = byte('_')
	}
}

// layers: all lower layers, lowest to highest
// chars: original input
// i: left index of this layer
// j: right index of this layer
// On an asymetric layer, i.e. chars[i] != chars[j], this pushes the left character onto layers first, then the right character
func updateLayers(layers []*layer, chars []byte, i, j int) ([]*layer, int, int) {
	var (
		top *layer
	)
	if len(layers) > 0 {
		top = layers[len(layers)-1]
	}
	asymetric := chars[i] != chars[j]
	if asymetric {
		lay1 := &layer{
			i:     i,
			j:     i,
			chars: chars,
		}
		lay2 := &layer{
			i:     j,
			j:     j,
			chars: chars,
		}
		top.push(lay1)
		lay1.push(lay2)
		layers = append(layers, lay1)
		layers = append(layers, lay2)
	} else {
		lay1 := &layer{
			i:     i,
			j:     j,
			chars: chars,
		}
		top.push(lay1)
		layers = append(layers, lay1)
	}
	i++
	j--
	return layers, i, j
}
