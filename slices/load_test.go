package slices

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestX(t *testing.T) {
	_ = assert.New(t)
	_ = require.New(t)
}

func TestLoadJson(t *testing.T) {

	txt := `[[
		true, false, true
	], [
		false, true, true
	], [
		true, true, false
]]`

	grid, err := LoadJson(strings.NewReader(txt))

	require.NoError(t, err)

	type coord struct {
		x, y uint
	}

	testCases := map[coord]bool{
		{0, 0}: true, {0, 1}: false, {0, 2}: true,
		{1, 0}: false, {1, 1}: true, {1, 2}: true,
		{2, 0}: true, {2, 1}: true, {2, 2}: false,
	}

	for c, b := range testCases {
		t.Run(fmt.Sprintf("%+v", c), func(t *testing.T) {
			var actual bool
			require.NotPanics(t, func() {
				actual = grid.Get(c.x, c.y)
			})
			assert.Equal(t, b, actual)
		})
	}
}
