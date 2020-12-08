package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01Part1(t *testing.T) {
	d := &Day01{
		expenses: []int{1721, 979, 366, 299, 675, 1456},
	}
	entryA, entryB := d.findTwoExpenseEntries()
	assert.Equal(t, 2020, entryA+entryB)
}

func TestDay01Part2(t *testing.T) {
	d := &Day01{
		expenses: []int{1721, 979, 366, 299, 675, 1456},
	}
	entryA, entryB, entryC := d.findThreeExpenseEntries()
	assert.Equal(t, 2020, entryA+entryB+entryC)
}
