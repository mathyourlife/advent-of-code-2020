package solver

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay14ExecV2(t *testing.T) {
	d := &Day14{
		memory: map[int]int{},
	}
	content := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	d.execV2(content)
	sum := 0
	for _, v := range d.memory {
		sum += v
	}
	assert.Equal(t, 208, sum)
}

func TestDay14WriteV2(t *testing.T) {
	d := &Day14{
		memory: map[int]int{},
	}
	set, _, floating := d.parseMask("000000000000000000000000000000X1001X")
	assert.Equal(t, 18, set)
	assert.Equal(t, []int{5, 0}, floating)
	d.writeV2(100, 42, set, floating)
	assert.Equal(t, map[int]int{26: 100, 27: 100, 58: 100, 59: 100}, d.memory)

	set, _, floating = d.parseMask("00000000000000000000000000000000X0XX")
	assert.Equal(t, 0, set)
	assert.Equal(t, []int{3, 1, 0}, floating)
	d.writeV2(1, 26, set, floating)
	sum := 0
	for _, v := range d.memory {
		sum += v
	}
	assert.Equal(t, 208, sum)
}

func TestDay14RunPrograms(t *testing.T) {
	d := &Day14{
		memory: map[int]int{},
	}
	d.execPrograms(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
	assert.Equal(t, map[int]int{7: 101, 8: 64}, d.memory)
}

func TestDay14ExecProgram(t *testing.T) {
	d := &Day14{
		memory: map[int]int{},
	}
	mask, instructions := d.parseProgram(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
	set, unset, _ := d.parseMask(mask)
	d.execProgram(instructions, set, unset)
	assert.Equal(t, map[int]int{7: 101, 8: 64}, d.memory)
}

func TestDay14ParseProgram(t *testing.T) {
	d := &Day14{}
	mask, instructions := d.parseProgram(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`)
	assert.Equal(t, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X", mask)
	assert.Equal(t, [][]int{{8, 11}, {7, 101}, {8, 0}}, instructions)
}

func TestDay14MaskValue(t *testing.T) {
	d := &Day14{}
	var set, unset, masked int
	set, unset, _ = d.parseMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	masked = d.maskValue(11, set, unset)
	assert.Equal(t, 73, masked)
	masked = d.maskValue(101, set, unset)
	assert.Equal(t, 101, masked)
	masked = d.maskValue(0, set, unset)
	assert.Equal(t, 64, masked)
}

func TestDay14ParseMask(t *testing.T) {
	d := &Day14{}
	var set, unset int
	var floating []int
	set, unset, floating = d.parseMask("111111111111111111111111111111111111")
	assert.Equal(t, int(math.Pow(2, 36))-1, set)
	assert.Equal(t, 0, unset)
	assert.Equal(t, []int{}, floating)
	set, unset, floating = d.parseMask("000000000000000000000000000000000000")
	assert.Equal(t, 0, set)
	assert.Equal(t, int(math.Pow(2, 36))-1, unset)
	assert.Equal(t, []int{}, floating)
	set, unset, floating = d.parseMask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X")
	assert.Equal(t, int(math.Pow(2, 6)), set)
	assert.Equal(t, int(math.Pow(2, 1)), unset)
	assert.Equal(t, []int{35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 5, 4, 3, 2, 0}, floating)
}
