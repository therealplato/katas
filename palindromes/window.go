package main

type win struct {
	chars []byte // source
	even  bool   // true: looking for even length strings
	l     int    // 0 = leftmost
	r     int    // 0 = rightmost
}

func (w win) L() int {
	return w.l
}
func (w win) R() int {
	return len(w.chars) - w.r
}

func (w win) size() int {
	return w.R() - w.L()
}

func (w win) val() []byte {
	return w.chars[w.L():w.R()]
}

// subWindows returns windows contained inside w and smaller than w; respecting even-ness
func (w win) subWindows() []win {
	var (
		i    int // window width iterator
		j    int // window origin iterator
		wins = make([]win, 0)
	)

	if w.even {
		for i = 2; i < w.size(); i += 2 {
			for j = w.L(); (j + i) <= w.R(); j++ {
				x := win{
					chars: w.chars,
					even:  true,
					l:     j,
					r:     j + i,
				}
				wins = append(wins, x)
			}
		}
	} else {
		for i = 1; i < w.size(); i += 2 {
			for j = w.L(); (j + i) <= w.R(); j++ {
				x := win{
					chars: w.chars,
					l:     j,
					r:     j + i,
				}
				wins = append(wins, x)
			}
		}
	}
	return wins
}
