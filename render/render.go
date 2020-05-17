package render

import (
	"fmt"
	"io"
	"strings"
)

var DefaultCharset = map[bool]byte{true: 'O', false: ' '}

type Renderable interface {
	Width() uint
	Height() uint
	Get(uint, uint) bool
}

type Frame struct {
	game         Renderable
	buffer       io.Writer
	characterSet map[bool]byte
}

func NewFrame(game Renderable, buffer io.Writer) Frame {
	return Frame{game, buffer, DefaultCharset}
}

func (f Frame) print(a ...interface{}) (int, error) {
	return fmt.Fprint(f.buffer, a...)
}

func (f Frame) printf(format string, a ...interface{}) (int, error) {
	return fmt.Fprintf(f.buffer, format, a...)
}

func (f Frame) println(a ...interface{}) (int, error) {
	return fmt.Fprintln(f.buffer, a...)
}

func (f Frame) Render() error {
	width := f.game.Width()

	// header
	header := func(corner, spacer byte) string {
		header := strings.Builder{}
		header.Grow(int(width + 2))
		header.WriteByte(corner)

		for i := uint(0); i < width; i++ {
			header.WriteByte(spacer)
		}

		header.WriteByte(corner)

		return header.String()
	}('+', '-')

	_, _ = f.println(header)

	// lines
	line := func(edge byte) string {
		line := strings.Builder{}
		line.Grow(2*int(width) + 3)
		line.WriteByte(edge)

		for i := uint(0); i < width; i++ {
			line.WriteString("%s")
		}

		line.WriteByte(edge)
		line.WriteByte('\n')

		return line.String()
	}('|')

	height := f.game.Height()
	lineVals := make([]interface{}, int(width))
	for j := uint(0); j < height; j++ {
		for i := uint(0); i < width; i++ {
			lineVals[i] = string(f.characterSet[f.game.Get(i, j)])
		}
		_, _ = f.printf(line, lineVals...)
	}

	// footer = header
	_, _ = f.println(header)

	return nil
}
