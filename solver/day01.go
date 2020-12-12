package solver

import (
	"fmt"
)

type Day01 struct{}

func NewDay01() Solver {
	return &Day01{}
}

func (d *Day01) SolvePart1(content string) {
	expenses := InputToInts(content)
	entryA, entryB := d.findTwoExpenseEntries(expenses)
	fmt.Printf("%d * %d = %d\n", entryA, entryB, entryA*entryB)
}

func (d *Day01) SolvePart2(content string) {
	expenses := InputToInts(content)
	entryA, entryB, entryC := d.findThreeExpenseEntries(expenses)
	fmt.Printf("%d * %d * %d = %d\n", entryA, entryB, entryC, entryA*entryB*entryC)
}

func (d *Day01) findTwoExpenseEntries(expenses []int) (int, int) {
	// brute force since it isn't a long list
	for i := 0; i < len(expenses)-1; i++ {
		for j := i + 1; j < len(expenses); j++ {
			if expenses[i]+expenses[j] == 2020 {
				return expenses[i], expenses[j]
			}
		}
	}
	return 0, 0
}

func (d *Day01) findThreeExpenseEntries(expenses []int) (int, int, int) {
	// even more brute force since it still isn't a long list
	for i := 0; i < len(expenses)-2; i++ {
		for j := i + 1; j < len(expenses); j++ {
			for k := j + 1; k < len(expenses); k++ {
				if expenses[i]+expenses[j]+expenses[k] == 2020 {
					return expenses[i], expenses[j], expenses[k]
				}
			}
		}
	}
	return 0, 0, 0
}
