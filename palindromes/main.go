package main

func main() {}

func pal(s string) bool {
	if len(s) < 2 {
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

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
			for j = w.L(); j < w.R(); j++ {
				x := win{
					chars: w.chars,
					l:     j,
					r:     j + i,
				}
				wins = append(wins, x)
			}
		}
	} else {
		for i = 1; i < w.size(); i += 2 {
			for j = w.L(); j < w.R(); j++ {
				x := win{
					chars: w.chars,
					l:     j,
					r:     j + i,
				}
				wins = append(wins, x)
			}
		}
	}
	//
	// if w.size() <= 1 {
	// 	return nil
	// }
	// if w.even {
	// 	j = 2
	// 	if w.size() <= 2 {
	// 		return nil
	// 	}
	// }
	// for i < w.R() {
	// 	x := win{
	// 		chars: w.chars,
	// 		even:  w.even,
	// 		l:     i,
	// 		r:     i + j,
	// 	}
	// 	wins = append(wins, x)
	// 	i++
	// }
	return wins
}

func subs(w win) [][]byte {
	subs := evenSubs(w)
	subs = append(subs, oddSubs(w)...)
	return subs
}

func evenSubs(w win) [][]byte {
	// L:=0: center between leftmost and second left char

	if w.size() == 2 {
		return [][]byte{
			w.val(),
		}
	}

	return nil
}
func oddSubs(w win) [][]byte {
	if w.even {
		return nil
	}
	// L=0: center at leftmost char
	if w.size() <= 2 {
		return [][]byte{
			[]byte{
				w.chars[w.L()],
			},
		}
	}
	subsubs := make([][]byte, 0)
	for _, ww := range w.subWindows() {
		subsubs = append(subsubs, subs(ww)...)
	}
	// tmp := win{
	// 	chars: w.chars,
	// 	even:  false,
	// }
	// // slide center point from left to right:
	// for i := w.L(); i < w.R(); i++ {
	// 	subsubs = append(subsubs, subs(tmp)...)
	// 	// collapse window 1 char per side:
	// 	tmp.l = w.l + 1
	// 	tmp.r = w.r + 1
	// 	// spaceLeft := i
	// 	// spaceRight =
	// }
	return subsubs
}
