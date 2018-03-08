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

func TestFindsSubstrings(t *testing.T) {

	type testcase struct {
		in win
		x  [][]byte
	}
	tcs := []testcase{
		testcase{
			in: win{
				chars: []byte("aa"),
				even:  true,
			},
			x: [][]byte{
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
		assert.Equal(t, tc.x, res)
	}
}

func toBBB(ss ...string) [][]byte {
	bbb := make([][]byte, len(ss))
	for _, s := range ss {
		bb := []byte(s)
		bbb = append(bbb, bb)
	}
	return bbb
}
