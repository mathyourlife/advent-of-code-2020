package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type Day02 struct{}

func NewDay02() Solver {
	d := &Day02{}
	return d
}

func (d *Day02) SolvePart1(content string) {
	records := strings.Split(content, "\n")
	numValid := d.countValidPasswordsV1(records)
	fmt.Printf("%d valid passwords by frequency\n", numValid)
}

func (d *Day02) SolvePart2(content string) {
	records := strings.Split(content, "\n")
	numValid := d.countValidPasswordsV2(records)
	fmt.Printf("%d valid passwords by position\n", numValid)
}

func (d *Day02) countValidPasswordsV1(records []string) int {
	count := 0
	for _, record := range records {
		min, max, letter, password := d.parseLine(record)
		if d.checkFrequency(min, max, letter, password) {
			count++
		}
	}
	return count
}

func (d *Day02) countValidPasswordsV2(records []string) int {
	count := 0
	for _, record := range records {
		min, max, letter, password := d.parseLine(record)
		if d.checkPosition(min, max, letter, password) {
			count++
		}
	}
	return count
}

func (d *Day02) parseLine(record string) (int, int, string, string) {
	parts := strings.Split(strings.TrimSpace(record), " ")
	freq := strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(freq[0])
	max, _ := strconv.Atoi(freq[1])
	return min, max, parts[1][0:1], parts[2]
}

func (d *Day02) checkFrequency(min, max int, letter, password string) bool {
	count := strings.Count(password, letter)
	return count >= min && count <= max
}

func (d *Day02) checkPosition(min, max int, letter, password string) bool {
	match := 0
	if password[min-1:min] == letter {
		match++
	}
	if password[max-1:max] == letter {
		match++
	}
	return match == 1
}
