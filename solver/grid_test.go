package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapStringGridGetLocation(t *testing.T) {
	g := NewWrapStringGrid(`ab
cd
ef
gh`)
	assert.Equal(t, "a", g.GetLocation(0, 0))
	assert.Equal(t, "d", g.GetLocation(1, 1))
	assert.Equal(t, "d", g.GetLocation(1, 3))
	assert.Equal(t, "h", g.GetLocation(3, 1))
}

func TestWrapStringGridDims(t *testing.T) {
	g := NewWrapStringGrid(`1234567
1234567
1234567
1234567`)
	assert.Equal(t, 4, g.rows)
	assert.Equal(t, 7, g.cols)
}
