package solver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTEMPLATE(t *testing.T) {
	d := &DayTEMPLATE{}
	assert.NotEqual(t, 0, d)
}
