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
		x  []byte
	}
	tcs := []testcase{{
		in: win{
			chars: []byte("aa"),
			even:  true,
		},
		x: []byte("aa"),
	}}

	for _, tc := range tcs {
		res := subs(tc.in)
		assert.Equal(t, tc.x, res)
	}
}
