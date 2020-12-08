package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay02checkFrequency(t *testing.T) {
	d := &Day02{}
	assert.True(t, d.checkFrequency(1, 3, "a", "abcde"))
	assert.False(t, d.checkFrequency(1, 3, "b", "cdefg"))
	assert.True(t, d.checkFrequency(2, 9, "c", "ccccccccc"))
}

func TestDay02checkPosition(t *testing.T) {
	d := &Day02{}
	assert.True(t, d.checkPosition(1, 3, "a", "abcde"))
	assert.False(t, d.checkPosition(1, 3, "b", "cdefg"))
	assert.False(t, d.checkPosition(2, 9, "c", "ccccccccc"))
}

func TestDay02Parse(t *testing.T) {
	d := &Day02{}
	min, max, letter, password := d.parseLine("1-3 b: cdefg")
	assert.Equal(t, 1, min)
	assert.Equal(t, 3, max)
	assert.Equal(t, "b", letter)
	assert.Equal(t, "cdefg", password)
}

func TestDay02Part1(t *testing.T) {
	d := &Day02{
		records: []string{
			"1-3 a: abcde",
			"1-3 b: cdefg",
			"2-9 c: ccccccccc",
		},
	}
	assert.Equal(t, 2, d.countValidPasswordsV1())
}
