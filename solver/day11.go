package solver

import (
	"fmt"
	"log"
	"strings"
)

type Day11 struct {
	mode string
	grid string
}

func NewDay11() Solver {
	return &Day11{}
}

func (d *Day11) SolvePart1(content string) {

	d.mode = "around"
	stabilized := d.stabilize(content)
	fmt.Printf("%d occupied seats\n", strings.Count(stabilized, "#"))
}

func (d *Day11) SolvePart2(content string) {

	d.mode = "seen"
	stabilized := d.stabilize(content)
	fmt.Printf("%d occupied seats\n", strings.Count(stabilized, "#"))
}

func (d *Day11) stabilize(grid string) string {
	var next string
	for {
		next = d.step(grid)
		if next == grid {
			break
		}
		grid = next
	}
	return next
}

func (d *Day11) step(grid string) string {
	next := ""

	var limit int
	switch d.mode {
	case "around":
		limit = 4
	case "seen":
		limit = 5
	}

	rows, cols := d.getDims(grid)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			loc := d.getLocation(grid, row, col, rows, cols)
			if loc == "." {
				// floor locations don't change
				next += "."
				continue
			}
			var occupied int
			switch d.mode {
			case "around":
				occupied = d.getOccupiedAround(grid, row, col, rows, cols)
			case "seen":
				occupied = d.getOccupiedSeen(grid, row, col, rows, cols)
			default:
				log.Fatalf("unrecognized mode '%s'", d.mode)
			}
			if loc == "L" && occupied == 0 {
				next += "#"
			} else if loc == "#" && occupied >= limit {
				next += "L"
			} else {
				next += loc
			}
		}
		if row < rows-1 {
			next += "\n"
		}
	}
	return next
}

func (d *Day11) getOccupiedSeen(grid string, row, col, rows, cols int) int {
	// log.Println("counting for", row, col)
	directions := [][]int{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}
	occupied := 0
	for _, direction := range directions {
		// log.Println("looking in", direction)
		p := []int{row, col}
		for {
			p[0] += direction[0]
			p[1] += direction[1]
			if p[0] < 0 || p[1] < 0 || p[0] >= rows || p[1] >= cols {
				break
			}
			// log.Println(p[0], p[1], d.getLocation(grid, p[0], p[1], rows, cols))
			loc := d.getLocation(grid, p[0], p[1], rows, cols)
			if loc == "#" {
				occupied++
				break
			} else if loc == "L" {
				break
			}
		}
	}
	return occupied
}

func (d *Day11) getOccupiedAround(grid string, row, col, rows, cols int) int {
	occupied := 0
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			// skip the middle or locations outside the grid
			if (x == 0 && y == 0) || row+y < 0 || row+y >= rows || col+x < 0 || col+x >= cols {
				continue
			}
			if d.getLocation(grid, row+y, col+x, rows, cols) == "#" {
				occupied++
			}
		}
	}
	return occupied
}

func (d *Day11) getLocation(grid string, row, col, rows, cols int) string {
	pos := (cols+1)*(row) + (col % cols)
	return grid[pos : pos+1]
}

func (d *Day11) getDims(grid string) (int, int) {
	rows := strings.Split(strings.TrimSpace(grid), "\n")
	return len(rows), len(rows[0])
}
