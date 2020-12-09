package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03CountTrees(t *testing.T) {
	d := &Day03{
		grid: `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`,
	}
	rows, cols := d.getDims()
	row, col := 0, 0
	countTrees := 0
	for {
		if row >= rows {
			break
		}
		loc := d.getLocation(row, col, rows, cols)
		if loc == "#" {
			countTrees++
		}
		col += 3
		row += 1
	}
	assert.Equal(t, 7, countTrees)
}

func TestDay03GetLocation(t *testing.T) {
	d := &Day03{
		grid: `ab
cd
ef
gh`,
	}
	width, height := d.getDims()
	assert.Equal(t, "a", d.getLocation(0, 0, width, height))
	assert.Equal(t, "d", d.getLocation(1, 1, width, height))
	assert.Equal(t, "d", d.getLocation(1, 3, width, height))
}

func TestDay03GetDims(t *testing.T) {
	d := &Day03{
		grid: `1234567
1234567
1234567
1234567`,
	}
	rows, cols := d.getDims()
	assert.Equal(t, 4, rows)
	assert.Equal(t, 7, cols)
}
