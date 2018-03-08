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
	// i := w.L()
	// for i; i < w.R(); i++ {
	// tmp := dupe(w)
	// subsubs := subs(tmp)
	// 	// spaceLeft := i
	// 	// spaceRight =
	// }

	return nil
}
func oddSubs(w win) [][]byte {
	// L=0: center at leftmost char
	if len(w.chars[w.L():w.R()]) == 1 {
		return [][]byte{
			[]byte{
				w.chars[w.L()],
			},
		}
	}
	return nil
}

func dupe(w win) win {
	return win{
		chars: w.chars,
		even:  w.even,
		l:     w.l,
		r:     w.r,
	}
}
