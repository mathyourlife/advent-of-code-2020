package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type Day07 struct{}

func NewDay07() Solver {
	d := &Day07{}
	return d
}

func (d *Day07) Reset() {}

func (d *Day07) SolvePart1(content string) {
	rs := NewDay07RuleSet(content)

	countGold := 0
	for _, rule := range rs.Rules {
		contains := rs.containsColors(rule.Color, []string{})
		for _, contain := range contains {
			if contain == "shiny gold" {
				countGold++
				break
			}
		}
	}
	fmt.Printf("%d colors can hold gold\n", countGold)
}

func (d *Day07) SolvePart2(content string) {
	rs := NewDay07RuleSet(content)

	totals := map[string]int{}
	for _, rule := range rs.Rules {
		if len(rule.Inner) == 0 {
			totals[rule.Color] = 0
		}
	}
	for {
		for _, rule := range rs.Rules {
			_, ok := totals[rule.Color]
			if ok {
				continue
			}
			knowAllInner := true
			total := 0
			for _, inner := range rule.Inner {
				t, ok := totals[inner.Color]
				if ok {
					total += inner.Count + inner.Count*t
				} else {
					knowAllInner = false
				}
			}
			if knowAllInner {
				totals[rule.Color] = total
			}
		}
		if len(totals) == len(rs.Rules) {
			break
		}
	}
	fmt.Printf("shiny gold holds %d bags\n", totals["shiny gold"])
}

type Day07RuleSet struct {
	Rules []*Day07Rule
}

func NewDay07RuleSet(content string) *Day07RuleSet {
	rs := &Day07RuleSet{
		Rules: []*Day07Rule{},
	}
	for _, rule := range strings.Split(content, "\n") {
		rs.Rules = append(rs.Rules, rs.parseRule(rule))
	}
	return rs
}

func (rs *Day07RuleSet) containsColors(color string, seenColors []string) []string {
	for _, rule := range rs.Rules {
		if rule.Color == color {
			for _, inner := range rule.Inner {
				found := false
				for _, seenColor := range seenColors {
					if seenColor == inner.Color {
						found = true
						break
					}
				}
				if !found {
					seenColors = append(seenColors, inner.Color)
					seenColors = rs.containsColors(inner.Color, seenColors)
				}
			}
		}
	}
	return seenColors
}

func (rs *Day07RuleSet) parseRule(rule string) *Day07Rule {
	r := &Day07Rule{}
	parts := strings.Split(rule, " bags contain ")
	r.Color = parts[0]
	contains := strings.Split(parts[1], ", ")
	for _, contain := range contains {
		innerSplit := strings.Split(contain, " ")
		inner := Day07Inner{}
		if innerSplit[0] == "no" {
			break
		}
		inner.Count, _ = strconv.Atoi(innerSplit[0])
		inner.Color = strings.Join(innerSplit[1:len(innerSplit)-1], " ")
		r.Inner = append(r.Inner, inner)
	}

	return r
}

type Day07Rule struct {
	Color string
	Inner []Day07Inner
}

type Day07Inner struct {
	Count int
	Color string
}

func (r Day07Rule) String() string {
	s := r.Color
	s += " bags contain "

	if len(r.Inner) == 0 {
		s += "no other bags."
		return s
	}
	contains := []string{}
	for _, inner := range r.Inner {
		bags := "bag"
		if inner.Count > 1 {
			bags = "bags"
		}
		contains = append(contains, fmt.Sprintf("%d %s %s", inner.Count, inner.Color, bags))
	}
	s += strings.Join(contains, ", ")
	s += "."

	return s
}
