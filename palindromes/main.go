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
	chars string // source
	even  bool   // true: looking for even length strings
	L     int    // 0 = leftmost
	R     int    // 0 = rightmost
}

func subs(w win) []string {
	// i=0: center at leftmost char
	// i=0: center between leftmost and second left char
	// i := 0
	// subs := evens()
	// subs = append(subs, odds()...)
	// for i; i < len(string); i++ {
	// 	spaceLeft := i
	// 	spaceRight =
	// }
	//
	return nil
}

func evenSubs(s win) []string {
	return nil
}
func oddSubs(s win) []string {
	return nil
}
