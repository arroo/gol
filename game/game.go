package game

import (
	"github.com/arroo/gol/render"
)

type Interface interface {
	render.Renderable
	Tick() Interface
}
