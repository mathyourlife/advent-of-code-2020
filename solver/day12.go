package solver

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Day12 struct {
}

func NewDay12() Solver {
	return &Day12{}
}

func (d *Day12) SolvePart1(content string) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	s.processCommands(strings.TrimSpace(content))

	fmt.Printf("state %v\n", s.x)
	fmt.Printf("manhattan distance %v\n", s.norm(s.x[0:2]))
}

func (d *Day12) SolvePart2(content string) {
	s := &Day12State{
		x: []int{0, 0, 0},
	}
	wp := &Day12State{
		x: []int{10, 1, 0},
	}
	s.processWPCommands(strings.TrimSpace(content), wp)

	fmt.Printf("ship position %v\n", s.x[0:2])
	fmt.Printf("waypoint position %v\n", wp.x[0:2])
	fmt.Printf("ship manhattan distance %v\n", s.norm(s.x[0:2]))
}

type Day12State struct {
	// 3 values contains position x, y and direction
	// where direction 0 = east, +1 rotates ccw, -1 rotates cw
	// 0 <= direction < 4
	x []int
}

// Take a rotation in degrees about the origin and
// return the associated 2x2 rotation matrix.
func (s *Day12State) R(deg int) []int {
	degf := float64(deg)
	return []int{
		int(math.Cos(degf * math.Pi / 180)), -int(math.Sin(degf * math.Pi / 180)),
		int(math.Sin(degf * math.Pi / 180)), int(math.Cos(degf * math.Pi / 180)),
	}
}

func (s *Day12State) norm(v []int) int {
	x := v[0]
	if x < 0 {
		x *= -1
	}
	y := v[1]
	if y < 0 {
		y *= -1
	}
	return x + y
}

func (s *Day12State) processWPCommands(cmds string, wp *Day12State) {
	for _, cmd := range strings.Split(cmds, "\n") {
		s.processWPCommand(cmd, wp)
	}
}

func (s *Day12State) processWPCommand(cmd string, wp *Day12State) {
	inc, _ := strconv.Atoi(cmd[1:])
	switch cmd[0:1] {
	case "E":
		wp.translate([]int{inc, 0})
	case "N":
		wp.translate([]int{0, inc})
	case "W":
		wp.translate([]int{-inc, 0})
	case "S":
		wp.translate([]int{0, -inc})
	case "R":
		R := s.R(-inc)
		new := []int{wp.x[0]*R[0] + wp.x[1]*R[1], wp.x[0]*R[2] + wp.x[1]*R[3]}
		wp.x[0] = new[0]
		wp.x[1] = new[1]
	case "L":
		R := s.R(inc)
		new := []int{wp.x[0]*R[0] + wp.x[1]*R[1], wp.x[0]*R[2] + wp.x[1]*R[3]}
		wp.x[0] = new[0]
		wp.x[1] = new[1]
	case "F":
		s.x[0] += inc * wp.x[0]
		s.x[1] += inc * wp.x[1]
	}
}

func (s *Day12State) processCommands(cmds string) {
	for _, cmd := range strings.Split(cmds, "\n") {
		s.processCommand(cmd)
	}
}

func (s *Day12State) processCommand(cmd string) {
	inc, _ := strconv.Atoi(cmd[1:])
	switch cmd[0:1] {
	case "E":
		s.translate([]int{inc, 0})
	case "N":
		s.translate([]int{0, inc})
	case "W":
		s.translate([]int{-inc, 0})
	case "S":
		s.translate([]int{0, -inc})
	case "R":
		s.rotate(-inc)
	case "L":
		s.rotate(inc)
	case "F":
		switch s.x[2] {
		case 0:
			s.translate([]int{inc, 0})
		case 1:
			s.translate([]int{0, inc})
		case 2:
			s.translate([]int{-inc, 0})
		case 3:
			s.translate([]int{0, -inc})
		}
	}
}

// horizontal and vertical transation
func (s *Day12State) translate(v []int) {
	s.x[0] += v[0]
	s.x[1] += v[1]
}

// turn in increments of 90 degrees with right hand rule around +z
// + is ccw about +z, - is cw about +z
func (s *Day12State) rotate(deg int) {
	inc := deg / 90
	s.x[2] += inc
	s.x[2] = s.x[2] % 4
	if s.x[2] < 0 {
		s.x[2] += 4
	}
}
