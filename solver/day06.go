package solver

import (
	"fmt"
	"strings"
)

type Day06 struct{}

func NewDay06() Solver {
	d := &Day06{}
	return d
}

func (d *Day06) SolvePart1(content string) {
	fullYesCount := 0
	for _, groupDeclarations := range strings.Split(content, "\n\n") {
		fullYesCount += d.countAnyYes(groupDeclarations)
	}
	fmt.Printf("total 'any' yes counts %d\n", fullYesCount)
}

func (d *Day06) SolvePart2(content string) {
	fullYesCount := 0
	for _, groupDeclarations := range strings.Split(content, "\n\n") {
		fullYesCount += d.countAllYes(groupDeclarations)
	}
	fmt.Printf("total 'all' yes counts %d\n", fullYesCount)
}

func (d *Day06) countAllYes(groupDeclarations string) int {
	declarations := strings.Split(groupDeclarations, "\n")
	union := map[string]bool{}
	for i := 0; i < len(declarations[0]); i++ {
		union[declarations[0][i:i+1]] = true
	}
	for _, declaration := range declarations[1:] {
		yes := map[string]bool{}
		for i := 0; i < len(declaration); i++ {
			yes[declaration[i:i+1]] = true
		}
		for question := range union {
			if !yes[question] {
				delete(union, question)
			}
		}
	}
	return len(union)
}

func (d *Day06) countAnyYes(groupDeclarations string) int {
	declarations := strings.ReplaceAll(groupDeclarations, "\n", "")
	yes := map[string]bool{}
	for i := 0; i < len(declarations); i++ {
		yes[declarations[i:i+1]] = true
	}
	return len(yes)
}
