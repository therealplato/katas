package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMain(t *testing.T) {
	t.Run("zero holes", func(t *testing.T) {
		t.Run("hill 0 1 0", func(t *testing.T) {
			c := Chart{0, 1, 0}
			d := fill(c)
			assert.Equal(t, 0, d)
		})
		t.Run("hill 0 9 5", func(t *testing.T) {
			c := Chart{0, 9, 5}
			d := fill(c)
			assert.Equal(t, 0, d)
		})
		t.Run("slope 54321", func(t *testing.T) {
			c := Chart{5, 4, 3, 2, 1}
			d := fill(c)
			assert.Equal(t, 0, d)
		})
	})
	t.Run("one hole", func(t *testing.T) {
		t.Run("valley 1 0 1", func(t *testing.T) {
			c := Chart{1, 0, 1}
			d := fill(c)
			assert.Equal(t, 1, d)
		})
		t.Run("valley 1 0 2", func(t *testing.T) {
			c := Chart{1, 0, 2}
			d := fill(c)
			assert.Equal(t, 1, d)
		})
		t.Run("valley 2 0 2", func(t *testing.T) {
			c := Chart{2, 0, 2}
			d := fill(c)
			assert.Equal(t, 2, d)
		})
		t.Run("valley 4 3 0 2", func(t *testing.T) {
			c := Chart{4, 3, 0, 2}
			d := fill(c)
			assert.Equal(t, 2, d)
		})
	})
	t.Run("two holes", func(t *testing.T) {
		t.Run("valleys 10101", func(t *testing.T) {
			c := Chart{1, 0, 1, 0, 1}
			d := fill(c)
			assert.Equal(t, 2, d)
		})
		t.Run("valleys 50423", func(t *testing.T) {
			c := Chart{5, 0, 4, 2, 3}
			d := fill(c)
			assert.Equal(t, 5, d)
		})
	})
}
