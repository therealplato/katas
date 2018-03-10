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
		assert.Equal(t, "__na__", actual)
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
	assert.Equal(t, 0, lay0.i)
	assert.Equal(t, 2, lay0.j)
	assert.Nil(t, lay0.down)
	require.NotNil(t, lay0.up)

	lay1 := lay0.up
	assert.Equal(t, lay1.chars, in)
	assert.Equal(t, 1, lay1.i)
	assert.Equal(t, 1, lay1.j)
	assert.NotNil(t, lay1.down)
	require.Nil(t, lay1.up)
}

func TestBuildsAsymetricLayers(t *testing.T) {
	in := []byte("ab")
	lay0 := build(in)
	assert.Equal(t, lay0.chars, in)
	assert.Equal(t, 0, lay0.i)
	assert.Equal(t, 0, lay0.j)
	assert.Nil(t, lay0.down)
	require.NotNil(t, lay0.up)

	lay1 := lay0.up
	assert.Equal(t, lay1.chars, in)
	assert.Equal(t, 1, lay1.i)
	assert.Equal(t, 1, lay1.j)
	assert.NotNil(t, lay1.down)
	require.Nil(t, lay1.up)
}
