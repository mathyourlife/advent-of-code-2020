package solver

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day04 struct{}

func NewDay04() Solver {
	d := &Day04{}
	return d
}

func (d *Day04) SolvePart1(content string) {
	fmt.Printf("%d accepted passports\n", d.countAcceptedPassports(
		strings.TrimSpace(content)))
}

func (d *Day04) SolvePart2(content string) {
	fmt.Printf("%d verified passports\n", d.verifiedAcceptedPassports(
		strings.TrimSpace(content)))
}

func (d *Day04) verifiedAcceptedPassports(content string) int {
	tally := 0
	for _, passport := range d.splitPassports(content) {
		if d.isVerified(passport) {
			tally++
		}
	}
	return tally
}

func (d *Day04) countAcceptedPassports(content string) int {
	tally := 0
	for _, passport := range d.splitPassports(content) {
		if d.isAccepted(passport) {
			tally++
		}
	}
	return tally
}

func (d *Day04) splitPassports(content string) []string {
	return strings.Split(content, "\n\n")
}

func (d *Day04) isVerified(passport string) bool {
	passport = strings.ReplaceAll(passport, "\n", " ")
	attrs := map[string]string{}
	for _, attr := range strings.Split(passport, " ") {
		parts := strings.Split(attr, ":")
		attrs[parts[0]] = parts[1]
	}
	// remove country id since we don't care
	delete(attrs, "cid")

	if len(attrs) != 7 {
		return false
	}

	byr, _ := strconv.Atoi(attrs["byr"])
	if byr < 1920 || byr > 2002 {
		return false
	}
	iyr, _ := strconv.Atoi(attrs["iyr"])
	if iyr < 2010 || iyr > 2020 {
		return false
	}
	eyr, _ := strconv.Atoi(attrs["eyr"])
	if eyr < 2020 || eyr > 2030 {
		return false
	}
	if len(attrs["pid"]) != 9 {
		return false
	}
	hgt := attrs["hgt"]
	switch hgt[len(hgt)-2:] {
	case "cm":
		hgtVal, _ := strconv.Atoi(hgt[:len(hgt)-2])
		if hgtVal < 150 || hgtVal > 193 {
			return false
		}
	case "in":
		hgtVal, _ := strconv.Atoi(hgt[:len(hgt)-2])
		if hgtVal < 59 || hgtVal > 76 {
			return false
		}
	default:
		return false
	}

	hcl := attrs["hcl"]
	if len(hcl) != 7 {
		return false
	}
	validHcl, _ := regexp.Match("^[0-9a-f]*$", []byte(hcl[1:]))
	if !validHcl {
		return false
	}

	validEcl := false
	for _, valid := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if valid == attrs["ecl"] {
			validEcl = true
			break
		}
	}
	if !validEcl {
		return false
	}

	return true
}

func (d *Day04) isAccepted(passport string) bool {
	passport = strings.ReplaceAll(passport, "\n", " ")
	attrs := map[string]bool{}
	for _, attr := range strings.Split(passport, " ") {
		key := strings.Split(attr, ":")[0]
		attrs[key] = true
	}
	// remove country id since we don't care
	delete(attrs, "cid")
	return len(attrs) == 7
}
