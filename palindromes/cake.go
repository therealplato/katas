package main

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
	// we'll destructively modify slice header:
	c := make([]byte, len(chars), len(chars))
	copy(c, chars)

	// layers[0] = bottom
	layers := make([]layer, 0)

	i := 0
	j := len(c) - 1
	for len(c) > 0 {
		lay := layer{
			i:     i,
			j:     j,
			chars: chars,
		}
		if len(layers) > 0 {
			lay.down = &(layers[len(layers)])
			(layers[len(layers)]).up = &lay
		}
		layers = append(layers, lay)
		if len(c) <= 1 {
			break
		}
		i++
		j--

		c = c[i:j]
	}
	if len(layers) == 0 {
		return nil
	}
	return &layers[0]
}

// func edges(c []byte) (l, r byte) {
// 	if len(c) == 1 {
// 		return c[0], c[0]
// 	}
// 	return c[0], c[len(c)]
// }

func spaceFill(bb []byte) {
	for i := range bb {
		bb[i] = byte(' ')
	}
}
