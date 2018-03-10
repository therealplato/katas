package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCakeFormatting(t *testing.T) {
	t.Run(`formats layer "a"`, func(t *testing.T) {
		bottom := build([]byte("a"))
		assert.Equal(t, "a", bottom.String())
	})
	t.Run(`formats layer "__na__"`, func(t *testing.T) {
		mid := layer{
			chars: []byte("panama"),
			i:     2,
			j:     3,
		}
		actual := mid.String()
		assert.Equal(t, "  na  ", actual)
	})
}

func TestBuildsOneLayer(t *testing.T) {
	in := []byte("a")
	lay := build(in)
	assert.Nil(t, lay.up)
	assert.Nil(t, lay.down)
	assert.Equal(t, lay.chars, in)
	assert.Equal(t, lay.i, 0)
	assert.Equal(t, lay.j, 0)
}

func TestBuildsTwoLayers(t *testing.T) {
	in := []byte("aba")
	lay0 := build(in)
	assert.Equal(t, lay0.chars, in)
	assert.Equal(t, lay0.i, 0)
	assert.Equal(t, lay0.j, 2)
	assert.Nil(t, lay0.down)
	require.NotNil(t, lay0.up)
}
