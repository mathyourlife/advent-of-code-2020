package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay09ValidateSequence(t *testing.T) {
	d := &Day09{}
	seq := []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
	valid, num, pos := d.validateSequence(seq, 5)
	assert.False(t, valid)
	assert.Equal(t, 127, num)
	assert.Equal(t, 14, pos)

	seq = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182}
	valid, num, pos = d.validateSequence(seq, 5)
	assert.True(t, valid)
	assert.Equal(t, 0, num)
	assert.Equal(t, 0, pos)
}

func TestDay09ValidateWindow(t *testing.T) {
	d := &Day09{}

	seq := []int{35, 20, 15, 25, 47}
	assert.True(t, d.validateWindow(seq, 40))

	seq = []int{95, 102, 117, 150, 182}
	assert.False(t, d.validateWindow(seq, 127))
}
