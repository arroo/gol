package slices

import (
	"github.com/arroo/gol/game" // imported only to satisfy the interface *cough* contracts *cough*
)

type grid [][]bool

func NewGrid(x, y uint) grid {
	g := make([][]bool, x)
	for i := range g {
		g[i] = make([]bool, y)
	}

	return g
}

func (g grid) Width() uint {
	return uint(len(g))
}

func (g grid) Height() uint {
	return uint(len(g[0]))
}

func (_ grid) left(x uint) bool {
	return x == 0
}

func (g grid) right(x uint) bool {
	return x == uint(len(g))-1
}

func (_ grid) top(y uint) bool {
	return y == 0
}

func (g grid) bottom(y uint) bool {
	return y == uint(len(g[0]))-1
}

func (g grid) Neighbours(x, y uint) []bool {

	type coord struct {
		x, y uint
	}

	coords := map[coord]bool{
		{x - 1, y - 1}: !g.left(x) && !g.top(y),
		{x - 1, y}:     !g.left(x),
		{x - 1, y + 1}: !g.left(x) && !g.bottom(y),
		{x, y - 1}:     !g.top(y),
		{x, y + 1}:     !g.bottom(y),
		{x + 1, y - 1}: !g.right(x) && !g.top(y),
		{x + 1, y}:     !g.right(x),
		{x + 1, y + 1}: !g.right(x) && !g.bottom(y),
	}

	out := make([]bool, 0)
	for c, f := range coords {
		if f {
			out = append(out, g[c.x][c.y])
		}
	}

	return out
}

func (g grid) AliveNeighbours(x, y uint) (out uint) {
	for _, alive := range g.Neighbours(x, y) {
		if alive {
			out++
		}
	}

	return out
}

func (g grid) CellTick(x, y uint) bool {

	n := g.AliveNeighbours(x, y)

	return n == 3 || n == 2 && g.Get(x, y)
}

func (g grid) Set(x, y uint, v bool) {
	g[x][y] = v
}

func (g grid) Get(x, y uint) bool {
	return g[x][y]
}

func (g grid) SetRow(x uint, row []bool) {
	if len(row) != len(g[x]) {
		panic("attempt to set col " + string(x) + " to mismatched length: " + string(len(row)))
	}

	g[x] = row
}

func (g grid) Tick() game.Interface {

	x := uint(len(g))
	y := uint(len(g[0]))

	chs := make([]chan []bool, x)
	for i := uint(0); i < x; i++ {

		ch := make(chan []bool)
		chs[i] = ch

		go func(ch chan []bool, i uint) {
			out := make([]bool, y)

			for j := uint(0); j < y; j++ {
				out[j] = g.CellTick(i, j)
			}

			ch <- out
			close(ch)
		}(ch, i)

	}

	out := NewGrid(x, y)
	for i, ch := range chs {
		out.SetRow(uint(i), <-ch)
	}

	return out
}
