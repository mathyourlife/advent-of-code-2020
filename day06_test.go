package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay06countAllYes(t *testing.T) {
	d := &Day06{}
	assert.Equal(t, 3, d.countAllYes(`abc`))
	assert.Equal(t, 0, d.countAllYes(`a
b
c`))
	assert.Equal(t, 1, d.countAllYes(`ab
ac`))
	assert.Equal(t, 1, d.countAllYes(`a
a
a
a`))
	assert.Equal(t, 1, d.countAllYes(`b`))
}

func TestDay06countAnyYes(t *testing.T) {
	d := &Day06{}
	assert.Equal(t, 3, d.countAnyYes(`abc`))
	assert.Equal(t, 3, d.countAnyYes(`a
b
c`))
	assert.Equal(t, 3, d.countAnyYes(`ab
ac`))
	assert.Equal(t, 1, d.countAnyYes(`a
a
a
a`))
	assert.Equal(t, 1, d.countAnyYes(`b`))
}
