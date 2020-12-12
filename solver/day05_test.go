package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay05SeatID(t *testing.T) {
	d := &Day05{}
	assert.Equal(t, 357, d.seatID("FBFBBFFRLR"))
}

func TestDay05BinarySpacePartitioning(t *testing.T) {
	d := &Day05{}
	assert.Equal(t, 127, d.binarySpacePartitioning("BBBBBBB"))
	assert.Equal(t, 0, d.binarySpacePartitioning("FFFFFFF"))
	assert.Equal(t, 70, d.binarySpacePartitioning("BFFFBBF"))
	assert.Equal(t, 14, d.binarySpacePartitioning("FFFBBBF"))
	assert.Equal(t, 102, d.binarySpacePartitioning("BBFFBBF"))
	assert.Equal(t, 7, d.binarySpacePartitioning("RRR"))
	assert.Equal(t, 4, d.binarySpacePartitioning("RLL"))
}
