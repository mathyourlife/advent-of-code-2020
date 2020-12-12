package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay12ProccessWPCommands(t *testing.T) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	wp := &Day12State{
		x: []int{10, 1, 0},
	}
	s.processWPCommands(`F10
N3
F7
R90
F11`, wp)
	assert.Equal(t, []int{214, -72}, s.x[0:2])
	assert.Equal(t, []int{4, -10}, wp.x[0:2])
}

func TestDay12Waypoint(t *testing.T) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	wp := &Day12State{
		x: []int{10, 1, 0},
	}

	s.processWPCommand("F10", wp)
	s.processWPCommand("N3", wp)
	s.processWPCommand("F7", wp)
	s.processWPCommand("R90", wp)
	s.processWPCommand("F11", wp)

	assert.Equal(t, []int{214, -72}, s.x[0:2])
	assert.Equal(t, []int{4, -10}, wp.x[0:2])
}

func TestDay12R(t *testing.T) {
	s := &Day12State{}
	assert.Equal(t, []int{0, -1, 1, 0}, s.R(90))
	assert.Equal(t, []int{0, 1, -1, 0}, s.R(-90))
	assert.Equal(t, []int{-1, 0, 0, -1}, s.R(180))
	assert.Equal(t, []int{1, 0, 0, 1}, s.R(0))
}

func TestDay12Norm(t *testing.T) {
	s := &Day12State{}
	assert.Equal(t, 25, s.norm([]int{17, 8}))
	assert.Equal(t, 25, s.norm([]int{-8, -17}))
}

func TestDay12ProccessCommands(t *testing.T) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	s.processCommands(`F10
N3
F7
R90
F11`)
	assert.Equal(t, []int{17, -8, 3}, s.x)
}

func TestDay12ProcessCommand(t *testing.T) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	s.processCommand("N3")
	assert.Equal(t, []int{0, 3, 0}, s.x)
	s.processCommand("F2")
	assert.Equal(t, []int{2, 3, 0}, s.x)
	s.processCommand("W1")
	assert.Equal(t, []int{1, 3, 0}, s.x)
	s.processCommand("S4")
	assert.Equal(t, []int{1, -1, 0}, s.x)
	s.processCommand("E4")
	assert.Equal(t, []int{5, -1, 0}, s.x)
	s.processCommand("L180")
	assert.Equal(t, []int{5, -1, 2}, s.x)
	s.processCommand("R270")
	assert.Equal(t, []int{5, -1, 3}, s.x)
}

func TestDay12Translation(t *testing.T) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	s.translate([]int{1, 4})
	assert.Equal(t, []int{1, 4, 0}, s.x)
	s.translate([]int{-7, 12})
	assert.Equal(t, []int{-6, 16, 0}, s.x)
}

func TestDay12StateRotate(t *testing.T) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	s.rotate(90)
	assert.Equal(t, 1, s.x[2])
	s.rotate(180)
	assert.Equal(t, 3, s.x[2])
	s.rotate(-360)
	assert.Equal(t, 3, s.x[2])
	s.rotate(90)
	assert.Equal(t, 0, s.x[2])
}
