package solver

import (
	"fmt"
)

type Day03 struct{}

func NewDay03() Solver {
	d := &Day03{}
	return d
}

func (d *Day03) SolvePart1(content string) {
	grid := NewWrapStringGrid(content)
	trees := d.countTrees(grid, 3, 1)
	fmt.Printf("%d trees encountered\n", trees)
}

func (d *Day03) SolvePart2(content string) {
	grid := NewWrapStringGrid(content)
	product := 1
	product *= d.countTrees(grid, 1, 1)
	product *= d.countTrees(grid, 3, 1)
	product *= d.countTrees(grid, 5, 1)
	product *= d.countTrees(grid, 7, 1)
	product *= d.countTrees(grid, 1, 2)
	fmt.Printf("%d trees product\n", product)
}

func (d *Day03) countTrees(grid *WrapStringGrid, over, down int) int {
	rows, _ := grid.Dims()
	row, col := 0, 0
	countTrees := 0
	for {
		if row >= rows {
			break
		}
		loc := grid.GetLocation(row, col)
		if loc == "#" {
			countTrees++
		}
		col += over
		row += down
	}
	return countTrees
}
