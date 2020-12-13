package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay13ConsecutiveSearch(t *testing.T) {
	d := &Day13{}
	_, busIDs, offsets := d.parseProblem(`0
17,x,13,19`)
	assert.Equal(t, 3417, d.consecutiveSearch(busIDs, offsets))

	_, busIDs, offsets = d.parseProblem(`0
67,7,59,61`)
	assert.Equal(t, 754018, d.consecutiveSearch(busIDs, offsets))

	_, busIDs, offsets = d.parseProblem(`0
67,x,7,59,61`)
	assert.Equal(t, 779210, d.consecutiveSearch(busIDs, offsets))

	_, busIDs, offsets = d.parseProblem(`0
67,7,x,59,61`)
	assert.Equal(t, 1261476, d.consecutiveSearch(busIDs, offsets))

	_, busIDs, offsets = d.parseProblem(`0
1789,37,47,1889`)
	assert.Equal(t, 1202161486, d.consecutiveSearch(busIDs, offsets))
}

func TestDay13NextDeparture(t *testing.T) {
	d := &Day13{}
	assert.Equal(t, 945, d.nextDeparture(939, 7))
	assert.Equal(t, 949, d.nextDeparture(939, 13))
	assert.Equal(t, 944, d.nextDeparture(939, 59))
	assert.Equal(t, 961, d.nextDeparture(939, 31))
	assert.Equal(t, 950, d.nextDeparture(939, 19))
}

func TestDay13ParseProblem(t *testing.T) {
	d := &Day13{}
	ts, busIDs, offsets := d.parseProblem(`939
7,13,x,x,59,x,31,19`)
	assert.Equal(t, 939, ts)
	assert.Equal(t, []int{7, 13, 59, 31, 19}, busIDs)
	assert.Equal(t, []int{0, 1, 4, 6, 7}, offsets)
}
