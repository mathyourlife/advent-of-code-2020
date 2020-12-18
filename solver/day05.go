package solver

import (
	"fmt"
	"sort"
	"strings"
)

type Day05 struct{}

func NewDay05() Solver {
	d := &Day05{}
	return d
}

func (d *Day05) Reset() {}

func (d *Day05) SolvePart1(content string) {
	backOfPlane := 0
	for _, pass := range strings.Split(content, "\n") {
		seatid := d.seatID(pass)
		if seatid > backOfPlane {
			backOfPlane = seatid
		}
	}
	fmt.Printf("highest seat id %d\n", backOfPlane)
}

func (d *Day05) SolvePart2(content string) {
	passes := strings.Split(content, "\n")
	occupied := make([]int, 0, len(passes))
	for _, pass := range passes {
		seatid := d.seatID(pass)
		occupied = append(occupied, seatid)
	}
	sort.Ints(occupied)
	for i, val := range occupied {
		if occupied[i+1] != val+1 {
			fmt.Printf("my seat %d\n", val+1)
			return
		}
	}
}

func (d *Day05) seatID(pass string) int {
	return d.binarySpacePartitioning(pass[:7])*8 + d.binarySpacePartitioning(pass[7:])
}

func (d *Day05) binarySpacePartitioning(partitioning string) int {
	val := 0
	for i := 0; i < len(partitioning); i++ {
		val = val << 1
		if partitioning[i:i+1] == "B" || partitioning[i:i+1] == "R" {
			val++
		}
	}
	return val
}
