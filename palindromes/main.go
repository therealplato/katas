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

func subs(w win) [][]byte {
	subs := evenSubs(w)
	subs = append(subs, oddSubs(w)...)
	return subs
}

func evenSubs(w win) [][]byte {
	// L:=0: center between leftmost and second left char

	if len(w.chars[w.L():w.R()]) == 2 {
		return [][]byte{
			w.chars[w.L():w.R()],
		}
	}

	return nil
}
func oddSubs(w win) [][]byte {
	if w.even {
		return nil
	}
	// L=0: center at leftmost char
	if len(w.chars[w.L():w.R()]) <= 2 {
		return [][]byte{
			[]byte{
				w.chars[w.L()],
			},
		}
	}
	subsubs := make([][]byte, 0)
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
