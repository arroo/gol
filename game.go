package main

import (
	"io"
	"os"
	"time"

	"github.com/arroo/gol/game"
	"github.com/arroo/gol/render"
	"github.com/arroo/gol/slices"
)

var generators map[string]game.Interface = make(map[string]game.Interface)

func init() {
	file := "glider.json"
	f, err := os.Open(file)
	if err != nil {
		panic("unable to open" + file + ":" + err.Error())
	}
	grid, err := slices.LoadJson(f)
	if err != nil {
		panic("unable to parse" + file + ":" + err.Error())
	}

	generators["glider"] = grid
}

func init() {
	grid := slices.NewGrid(100, 52)

	baseX := uint(80)
	baseY := uint(26)

	grid.Set(baseX, baseY, true)
	grid.Set(baseX+1, baseY, true)
	grid.Set(baseX+4, baseY, true)
	grid.Set(baseX+5, baseY, true)
	grid.Set(baseX+6, baseY, true)
	grid.Set(baseX+1, baseY-2, true)
	grid.Set(baseX+3, baseY-1, true)

	generators["acorn"] = grid
}

func Clear(w io.Writer) {
	clear := "\033[H\033[2J"
	w.Write([]byte(clear))
}

func main() {

	g := "acorn"

	grid := generators[g]

	w := os.Stdout

	for {

		_ = render.NewFrame(grid, w).Render()
		grid = grid.Tick()
		time.Sleep(100 * time.Millisecond)
		Clear(w)
	}
}
