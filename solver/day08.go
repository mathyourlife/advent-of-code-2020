package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type Day08 struct {
	instructions []string
	accumulator  int
	index        int
}

func NewDay08() Solver {
	return &Day08{}
}

func (d *Day08) Reset() {
	d.instructions = []string{}
	d.accumulator = 0
	d.index = 0
}

func (d *Day08) SolvePart1(content string) {
	d.instructions = strings.Split(content, "\n")

	executed := make([]bool, len(d.instructions))
	for {
		executed[d.index] = true
		d.exec()
		if executed[d.index] {
			break
		}
	}
	fmt.Printf("found infinite loop with accumulator at %d\n", d.accumulator)
}

func (d *Day08) SolvePart2(content string) {

	original := strings.Split(content, "\n")
	for step := 0; step < len(original); step++ {
		d.instructions = make([]string, len(original))
		copy(d.instructions, original)
		d.accumulator = 0
		d.index = 0

		switch d.instructions[step][:3] {
		case "acc":
			continue
		case "nop":
			d.instructions[step] = "jmp" + d.instructions[step][3:]
		case "jmp":
			d.instructions[step] = "nop" + d.instructions[step][3:]
		}
		executed := make([]bool, len(d.instructions))
		for {
			executed[d.index] = true
			d.exec()
			if d.index >= len(d.instructions) {
				fmt.Printf("found it. Instruction %d (%s) is corrupt. Accumulator exits with a value of %d\n", step, original[step], d.accumulator)
				return
			}
			if executed[d.index] {
				break
			}
		}
	}
}

func (d *Day08) exec() {
	instruction := d.instructions[d.index]

	switch instruction[:3] {
	case "nop":
		d.index++
	case "acc":
		i, _ := strconv.Atoi(instruction[4:])
		d.accumulator += i
		d.index++
	case "jmp":
		i, _ := strconv.Atoi(instruction[4:])
		d.index += i
	}
}
