package main

func subs(w win) [][]byte {
	if w.size()%2 == 0 {
		return evenSubs(w)
	}
	return oddSubs(w)
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
