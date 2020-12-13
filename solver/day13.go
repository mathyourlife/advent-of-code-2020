package solver

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day13 struct {
}

func NewDay13() Solver {
	return &Day13{}
}

func (d *Day13) SolvePart1(content string) {
	ts, busIDs, _ := d.parseProblem(content)
	earliestTS := d.nextDeparture(ts, busIDs[0])
	earliestBus := busIDs[0]
	for _, busID := range busIDs[1:] {
		depart := d.nextDeparture(ts, busID)
		if earliestTS > depart {
			earliestTS = depart
			earliestBus = busID
		}
	}
	fmt.Printf("Bus %d will depart at %d which is in %d minutes\n", earliestBus, earliestTS, earliestTS-ts)
	fmt.Printf("code: %d\n", earliestBus*(earliestTS-ts))
}

func (d *Day13) SolvePart2(content string) {
	_, busIDs, offsets := d.parseProblem(content)
	ts := d.consecutiveSearch(busIDs, offsets)
	fmt.Printf("busses will arrive in consecutive minutes starting at ts = %d\n", ts)
	fmt.Printf("which is about %d years from when the busses started their routes\n", ts/60/24/365)
}

func (d *Day13) consecutiveSearch(busIDs, offsets []int) int {

	count := len(busIDs)

	// sort busIDs and offsets in descending order to make the search results
	// faster since products will grow faster
	sortedBusIDs := make([]int, count)
	copy(sortedBusIDs, busIDs)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedBusIDs)))

	sortedOffsets := make([]int, count)
	for sortedIdx, sortedBusID := range sortedBusIDs {
		for unsortedIdx, busID := range busIDs {
			if sortedBusID == busID {
				sortedOffsets[sortedIdx] = offsets[unsortedIdx]
			}
		}
	}

	ts := 0
	// start with ts 0 and increment by 1 to find first mod = 0
	// fix first offset at 16
	// then increment by 19's
	// waid for second mod to sync at 187
	// then increment by 19*17
	inc := 1
	lockedIn := 0
	for {
		// fmt.Printf("%5d", ts)
		mod := (ts + sortedOffsets[lockedIn]) % sortedBusIDs[lockedIn]
		// fmt.Printf("    +%d mod %3d = %3d (%2d)\n", sortedOffsets[lockedIn], sortedBusIDs[lockedIn], mod, sortedBusIDs[lockedIn]-mod)
		if mod == 0 {
			// locked in the next bus rotation, increase search by a product of its ID
			// fmt.Println("locked", sortedBusIDs[lockedIn])
			inc *= sortedBusIDs[lockedIn]
			lockedIn++
		}
		if lockedIn >= count {
			return ts
		}
		ts += inc
	}
	return 0
}

func (d *Day13) nextDeparture(ts, busID int) int {
	if ts%busID == 0 {
		return ts
	}
	w := ts / busID
	return (w + 1) * busID
}

func (d *Day13) parseProblem(content string) (int, []int, []int) {
	parts := strings.Split(content, "\n")
	ts, _ := strconv.Atoi(parts[0])
	busIDs := []int{}
	offsets := []int{}
	for offset, busStr := range strings.Split(parts[1], ",") {
		if busStr == "x" {
			continue
		}
		busID, _ := strconv.Atoi(busStr)
		busIDs = append(busIDs, busID)
		offsets = append(offsets, offset)
	}
	return ts, busIDs, offsets
}
