package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay07ContainsGold(t *testing.T) {
	rs := NewDay07RuleSet(`light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`)

	assert.Contains(t, rs.containsColors("bright white", []string{}), "shiny gold")
	assert.Contains(t, rs.containsColors("muted yellow", []string{}), "shiny gold")
	assert.Contains(t, rs.containsColors("dark orange", []string{}), "shiny gold")
	assert.Contains(t, rs.containsColors("light red", []string{}), "shiny gold")
	assert.NotContains(t, rs.containsColors("dark olive", []string{}), "shiny gold")
}

func TestDay07ContainsColors(t *testing.T) {
	rs := NewDay07RuleSet(`muted yellow bags contain 1 bright white bag, 2 light red bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
light red bags contain 3 shiny gold bags.`)

	assert.Equal(t, []string{"bright white", "muted yellow", "light red", "shiny gold"}, rs.containsColors("dark orange", []string{}))
}

func TestDay07ParseRule(t *testing.T) {
	rs := &Day07RuleSet{}

	rules := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`
	for _, rule := range strings.Split(rules, "\n") {
		assert.Equal(t, rule, rs.parseRule(rule).String())
	}
}
