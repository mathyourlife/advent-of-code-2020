package solver

import (
	"fmt"
	"sort"
)

type Day10 struct {
}

func NewDay10() Solver {
	return &Day10{}
}

func (d *Day10) Reset() {}

func (d *Day10) SolvePart1(content string) {
	fmt.Printf("just used excel\n")
}

func (d *Day10) SolvePart2(content string) {

	vals := []int{0}
	vals = append(vals, InputToInts(content)...)
	paths := make([]int, len(vals))
	sort.Ints(vals)
	paths[0] = 1

	for i := 1; i < len(vals); i++ {
		for j := i - 1; j >= 0; j-- {
			if vals[j]+3 < vals[i] {
				break
			}
			paths[i] += paths[j]
		}
	}
	fmt.Println("paths", paths)
	fmt.Printf("%d combinations of adapters\n", paths[len(paths)-1])

}
