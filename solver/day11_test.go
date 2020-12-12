package solver

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay11AroundStability(t *testing.T) {
	d := &Day11{
		mode: "around",
	}
	grid := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	assert.Equal(t, `#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`, d.stabilize(grid))

	assert.Equal(t, 37, strings.Count(d.stabilize(grid), "#"))

}

func TestDay11SeenStability(t *testing.T) {
	d := &Day11{
		mode: "seen",
	}
	grid := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	assert.Equal(t, `#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`, d.stabilize(grid))

	assert.Equal(t, 26, strings.Count(d.stabilize(grid), "#"))

}

func TestDay11Step(t *testing.T) {
	d := &Day11{
		mode: "around",
	}

	assert.Equal(t, `...
.#.
...`, d.step(`...
.L.
...`))

	assert.Equal(t, `#.#
.LL
.##`, d.step(`#.#
.##
.##`))

	states := []string{`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`,
		`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,
		`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`,
		`#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`,
		`#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##`,
		`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`}

	for i := 0; i < len(states)-1; i++ {
		assert.Equal(t, states[i+1], d.step(states[i]))
	}

}

func TestDay11GetOccupiedSeen(t *testing.T) {
	d := &Day11{}

	grid := `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`
	width, height := d.getDims(grid)
	assert.Equal(t, 8, d.getOccupiedSeen(grid, 4, 3, width, height))

	grid = `.............
.L.L.#.#.#.#.
.............`
	width, height = d.getDims(grid)
	assert.Equal(t, 0, d.getOccupiedSeen(grid, 1, 1, width, height))

	grid = `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`
	width, height = d.getDims(grid)
	assert.Equal(t, 0, d.getOccupiedSeen(grid, 3, 3, width, height))
}

func TestDay11GetOccupiedAround(t *testing.T) {
	d := &Day11{}

	grid := `##
..`
	width, height := d.getDims(grid)
	assert.Equal(t, 1, d.getOccupiedAround(grid, 0, 0, width, height))

	grid = `.....#..
.##.....
...##.##
...#....`

	width, height = d.getDims(grid)
	assert.Equal(t, 3, d.getOccupiedAround(grid, 2, 3, width, height))
}

func TestDay11GetLocation(t *testing.T) {
	d := &Day11{}
	grid := `ab
cd
ef
gh`
	width, height := d.getDims(grid)
	assert.Equal(t, "a", d.getLocation(grid, 0, 0, width, height))
	assert.Equal(t, "d", d.getLocation(grid, 1, 1, width, height))
	assert.Equal(t, "d", d.getLocation(grid, 1, 3, width, height))
}

func TestDay11GetDims(t *testing.T) {
	d := &Day11{}
	grid := `1234567
1234567
1234567
1234567`
	rows, cols := d.getDims(grid)
	assert.Equal(t, 4, rows)
	assert.Equal(t, 7, cols)
}
