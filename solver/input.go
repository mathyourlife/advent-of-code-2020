package solver

import (
	"log"
	"strconv"
	"strings"
)

func InputToInts(content string) []int {
	vals := []int{}
	for _, entry := range strings.Split(strings.TrimSpace(content), "\n") {
		val, err := strconv.Atoi(entry)
		if err != nil {
			log.Fatalf("unable to convert '%s' to an integer: %s", entry, err)
		}
		vals = append(vals, val)
	}
	return vals
}
