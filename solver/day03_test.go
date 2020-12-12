package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay03CountTreeProduct(t *testing.T) {
	d := &Day03{}

	grid := NewWrapStringGrid(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	product := 1
	product *= d.countTrees(grid, 1, 1)
	product *= d.countTrees(grid, 3, 1)
	product *= d.countTrees(grid, 5, 1)
	product *= d.countTrees(grid, 7, 1)
	product *= d.countTrees(grid, 1, 2)
	assert.Equal(t, 336, product)
}

func TestDay03CountTrees(t *testing.T) {
	d := &Day03{}

	grid := NewWrapStringGrid(`..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`)

	assert.Equal(t, 7, d.countTrees(grid, 3, 1))
}
