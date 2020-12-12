package solver

import "strings"

type WrapStringGrid struct {
	rows int
	cols int
	data string
}

func NewWrapStringGrid(data string) *WrapStringGrid {
	tmp := strings.Split(strings.TrimSpace(data), "\n")
	g := &WrapStringGrid{
		rows: len(tmp),
		cols: len(tmp[0]),
		data: strings.ReplaceAll(data, "\n", ""),
	}
	return g
}

func (g *WrapStringGrid) Dims() (rows int, cols int) {
	return g.rows, g.cols
}

func (g *WrapStringGrid) GetLocation(row, col int) string {
	pos := (g.cols)*(row) + (col % g.cols)
	return g.data[pos : pos+1]
}
