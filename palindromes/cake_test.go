package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLayerFormatting(t *testing.T) {
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
func TestCakeFormatting(t *testing.T) {
	t.Run(`"a" cake`, func(t *testing.T) {
		bottom := bake([]byte("a"))
		assert.Equal(t, "a\n", bottom.String())
	})
	t.Run(`"ab" cake`, func(t *testing.T) {
		bottom := bake([]byte("ab"))
		assert.Equal(t, `_b
a_
`, bottom.String())
	})
	t.Run(`"aba" cake`, func(t *testing.T) {
		bottom := bake([]byte("aba"))
		assert.Equal(t, `_b_
a_a
`, bottom.String())
	})
}

func TestBuildsOneLayer(t *testing.T) {
	in := []byte("a")
	cake := bake(in)
	lay := cake.bottom()
	assert.Nil(t, lay.up)
	assert.Nil(t, lay.down)
	assert.Equal(t, lay.chars, in)
	assert.Equal(t, lay.i, 0)
	assert.Equal(t, lay.j, 0)
}

func TestBuildsTwoLayers(t *testing.T) {
	in := []byte("aba")
	cake := bake(in)
	lay0 := cake.bottom()
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
	cake := bake(in)
	lay0 := cake.bottom()
	assert.Equal(t, lay0.chars, in)
	assert.Equal(t, 0, lay0.i)
	assert.Equal(t, 0, lay0.j)
	assert.Nil(t, lay0.down)
	require.NotNil(t, lay0.up)

	lay1 := lay0.up
	assert.Equal(t, lay1, cake.top())
	assert.Equal(t, lay1.chars, in)
	assert.Equal(t, 1, lay1.i)
	assert.Equal(t, 1, lay1.j)
	assert.NotNil(t, lay1.down)
	require.Nil(t, lay1.up)

}
