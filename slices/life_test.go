package slices

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid(10, 11)

	require := require.New(t)
	require.NotNil(g)

	// verify x
	require.Equal(10, len(g))

	// verify each y
	assert := assert.New(t)

	for i := range g {
		assert.NotNil(g[i])
		assert.Equal(11, len(g[i]))
	}
}

func TestNeighbours(t *testing.T) {

	t.Run("correct count & no panics", func(t *testing.T) {

		box := NewGrid(3, 3)
		col := NewGrid(1, 3)
		row := NewGrid(3, 1)

		testCases := map[string]struct {
			g    grid
			x, y uint
			c    int
		}{
			"box - middle":    {box, 1, 1, 8},
			"box - corner":    {box, 0, 0, 3},
			"box - top edge":  {box, 0, 1, 5},
			"box - side edge": {box, 1, 0, 5},
			"col - top":       {col, 0, 0, 1},
			"col - mid":       {col, 0, 1, 2},
			"col - bot":       {col, 0, 2, 1},
			"row - left":      {row, 0, 0, 1},
			"row - mid":       {row, 1, 0, 2},
			"row - right":     {row, 2, 0, 1},
		}

		for name, tc := range testCases {
			t.Run(name, func(t *testing.T) {
				// setup
				//g := NewGrid(3, 3)

				// exercise
				var n []bool
				require.NotPanics(t, func() {
					n = tc.g.Neighbours(tc.x, tc.y)
				})

				// verification
				assert.Equalf(t, tc.c, len(n), "count of neighbours")
			})
		}
	})
}

func TestNeighbourCount(t *testing.T) {

	box := grid{
		{false, true, false},
		{true, false, true},
		{false, false, true},
	}

	testCases := map[string]struct {
		g       grid
		x, y, c uint
	}{
		"something": {box, 1, 1, 4},
		"1":         {box, 0, 0, 2},
		"2":         {box, 2, 2, 1},
		"3":         {box, 0, 2, 2},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			// setup
			//g := NewGrid(3, 3)

			// exercise
			n := tc.g.AliveNeighbours(tc.x, tc.y)

			// verification
			assert.Equalf(t, tc.c, n, "alive neighbours")
		})
	}
}

func TestCellTick(t *testing.T) {
	box := grid{
		{false, true, false},
		{true, false, true},
		{false, false, true},
	}

	testCases := map[string]struct {
		x, y uint
		s    bool
	}{
		"live stays live": {0, 1, true},
		"dead stays dead": {0, 2, false},
		"live dies":       {1, 0, false},
		"dead lives":      {1, 2, true},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.s, box.CellTick(tc.x, tc.y))
		})
	}
}

func TestTick(t *testing.T) {
	testCases := []struct {
		seed, tick grid
	}{{
		seed: grid{
			{false, true, false},
			{true, false, true},
			{false, false, true},
		},
		tick: grid{
			{false, true, false},
			{false, false, true},
			{false, true, false},
		},
	}}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tc.tick, tc.seed.Tick())
		})
	}
}
