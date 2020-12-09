package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03CountTreeProduct(t *testing.T) {
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
	product := 1
	product *= d.countTrees(1, 1)
	product *= d.countTrees(3, 1)
	product *= d.countTrees(5, 1)
	product *= d.countTrees(7, 1)
	product *= d.countTrees(1, 2)
	assert.Equal(t, 336, product)
}

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
	assert.Equal(t, 7, d.countTrees(3, 1))
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
