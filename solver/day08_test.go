package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay08Exec(t *testing.T) {
	d := &Day08{
		instructions: []string{"nop +0", "acc -7"},
		index:        1,
		accumulator:  10,
	}
	d.exec()
	assert.Equal(t, 2, d.index)
	assert.Equal(t, 3, d.accumulator)

	d = &Day08{
		instructions: []string{"nop +0", "jmp -7"},
		index:        1,
		accumulator:  10,
	}
	d.exec()
	assert.Equal(t, -6, d.index)
	assert.Equal(t, 10, d.accumulator)

	d = &Day08{
		instructions: []string{"nop +0", "nop +10"},
		index:        1,
		accumulator:  10,
	}
	d.exec()
	assert.Equal(t, 2, d.index)
	assert.Equal(t, 10, d.accumulator)
}
