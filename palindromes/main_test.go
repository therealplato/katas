package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	yess := []string{
		"a",
		"aa",
		"aaa",
		"aba",
	}
	nos := []string{
		"ab",
		"abb",
		"abc",
	}

	for _, tc := range yess {
		res := pal(tc)
		assert.True(t, res, tc)
	}
	for _, tc := range nos {
		res := pal(tc)
		assert.False(t, res, tc)
	}
}

func TestSubWindows(t *testing.T) {
	type testcase struct {
		name string
		in   win
		out  []win
	}
	tcs := []testcase{
		testcase{
			name: "even window one char",
			in: win{
				chars: []byte("a"),
				even:  true,
			},
			out: nil,
		},
		testcase{
			name: "odd window one char",
			in: win{
				chars: []byte("a"),
				even:  false,
			},
			out: nil,
		},
		testcase{
			name: "even window two chars",
			in: win{
				chars: []byte("aa"),
				even:  true,
			},
			out: nil,
		},
		testcase{
			name: "odd window two chars",
			in: win{
				chars: []byte("aa"),
			},
			out: []win{
				win{
					chars: []byte("aa"),
					l:     0,
					r:     1,
				},
				win{
					chars: []byte("aa"),
					l:     1,
					r:     2,
				},
			},
		},
		testcase{
			name: "odd window three chars",
			in: win{
				chars: []byte("aaa"),
			},
			out: []win{
				win{
					chars: []byte("aaa"),
					l:     0,
					r:     1,
				},
				win{
					chars: []byte("aaa"),
					l:     1,
					r:     2,
				},
				win{
					chars: []byte("aaa"),
					l:     2,
					r:     3,
				},
			},
		},
		testcase{
			name: "odd window four chars",
			in: win{
				chars: []byte("aaaa"),
			},
			out: []win{
				win{
					chars: []byte("aaaa"),
					l:     0,
					r:     1,
				},
				win{
					chars: []byte("aaaa"),
					l:     1,
					r:     2,
				},
				win{
					chars: []byte("aaaa"),
					l:     2,
					r:     3,
				},
				win{
					chars: []byte("aaaa"),
					l:     0,
					r:     3,
				},
				win{
					chars: []byte("aaaa"),
					l:     1,
					r:     4,
				},
			},
		},
		testcase{
			name: "even window five chars",
			in: win{
				chars: []byte("aaaaa"),
				even:  true,
			},
			out: []win{
				win{
					chars: []byte("aaaaa"),
					even:  true,
					l:     0,
					r:     2,
				},
				win{
					chars: []byte("aaaaa"),
					even:  true,
					l:     1,
					r:     3,
				},
				win{
					chars: []byte("aaaaa"),
					even:  true,
					l:     2,
					r:     4,
				},
				win{
					chars: []byte("aaaaa"),
					even:  true,
					l:     3,
					r:     5,
				},
				win{
					chars: []byte("aaaaa"),
					even:  true,
					l:     0,
					r:     4,
				},
				win{
					chars: []byte("aaaaa"),
					even:  true,
					l:     1,
					r:     5,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			subs := tc.in.subWindows()
			for _, x := range tc.out {
				assert.Contains(t, subs, x, tc.name)
			}
		})
	}
}

func TestOddSubs(t *testing.T) {
	t.Run("one char input", func(t *testing.T) {
		res := oddSubs(win{
			chars: []byte("a"),
		})
		x := [][]byte{
			[]byte("a"),
		}
		assert.Equal(t, x, res)
	})

	t.Run("two char input", func(t *testing.T) {
		res := oddSubs(win{
			chars: []byte("aa"),
		})
		x := [][]byte{
			[]byte("a"),
		}
		assert.Equal(t, x, res)
	})

}

func TestFindsSubstrings(t *testing.T) {

	type testcase struct {
		in win
		x  [][]byte
	}
	tcs := []testcase{
		testcase{
			in: win{
				chars: []byte("a"),
			},
			x: [][]byte{
				[]byte("a"),
			},
		},
		testcase{
			in: win{
				chars: []byte("aa"),
			},
			x: [][]byte{
				[]byte("a"),
				[]byte("aa"),
			},
		},
		testcase{
			in: win{
				chars: []byte("aa"),
			},
			x: [][]byte{
				[]byte("a"),
				[]byte("aa"),
			},
		},
		testcase{
			in: win{
				chars: []byte("abba"),
				even:  true,
			},
			x: toBBB("aa", "bb", "aba", "abba"),
		},
	}

	for _, tc := range tcs {
		res := subs(tc.in)
		for _, xSub := range tc.x {
			assert.Contains(t, res, xSub, string(tc.in.chars))
		}
	}
}

func toBBB(ss ...string) [][]byte {
	bbb := make([][]byte, 0)
	for _, s := range ss {
		bb := []byte(s)
		bbb = append(bbb, bb)
	}
	return bbb
}
