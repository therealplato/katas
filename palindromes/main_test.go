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

	t.Run("three char input", func(t *testing.T) {
		res := oddSubs(win{
			chars: []byte("aaa"),
		})
		x := [][]byte{
			[]byte("a"),
			[]byte("aaa"),
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
