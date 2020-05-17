package slices

import (
	"encoding/json"
	"io"
)

func LoadJson(source io.Reader) (grid, error) {

	dec := json.NewDecoder(source)

	g := grid{}
	err := dec.Decode(&g)

	return g, err
}
