package solver

import (
	"fmt"
	"log"
	"sort"
)

type Day09 struct {
}

func NewDay09() Solver {
	return &Day09{}
}

func (d *Day09) Reset() {}

func (d *Day09) SolvePart1(content string) {
	seq := InputToInts(content)
	valid, target, pos := d.validateSequence(seq, 25)
	if valid {
		log.Println("it checks out")
	} else {
		log.Printf("%d at position %d doesn't conform to the rules\n", target, pos)
	}
}

func (d *Day09) SolvePart2(content string) {
	seq := InputToInts(content)

	valid, target, pos := d.validateSequence(seq, 25)
	if valid {
		log.Println("it checks out")
	}

	log.Printf("checking for window that sums to %d at position %d", target, pos)

	for scroll := pos - 1; scroll >= pos-400; scroll-- {
		index := scroll
		sum := seq[index]
		for {
			if sum > target {
				break
			}
			if sum == target {
				window := seq[index : scroll+1]
				fmt.Printf("found it %#v\n", seq[index:scroll+1])
				// note 'window' isn't a copy so the sort here modifies Day09Content
				sort.Ints(window)
				fmt.Printf("sum of max and min is %d + %d = %d\n", window[0], window[len(window)-1], window[0]+window[len(window)-1])
				return
			}
			index--
			sum += seq[index]
		}
	}
	fmt.Println("couldn't find it :(")
}

func (d *Day09) validateSequence(seq []int, preamble int) (bool, int, int) {
	window := make([]int, preamble)
	for i := 0; i < preamble; i++ {
		window[i] = seq[i]
	}
	ring := 0
	for i := preamble; i < len(seq); i++ {
		// log.Printf("checking %d, to place in index %d", seq[i], ring)
		// log.Println(window)
		if !d.validateWindow(window, seq[i]) {
			return false, seq[i], i
		}
		window[ring] = seq[i]
		ring++
		if ring >= preamble {
			ring = 0
		}
	}
	return true, 0, 0
}

func (d *Day09) validateWindow(window []int, target int) bool {
	for i := 0; i < len(window)-1; i++ {
		for j := i + 1; j < len(window); j++ {
			if window[i]+window[j] == target {
				return true
			}
		}
	}
	return false
}
