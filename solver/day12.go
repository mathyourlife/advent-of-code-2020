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
	s.processCommands(Day12Content)

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
	s.processWPCommands(Day12Content, wp)

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

const Day12Content = `W5
R90
W3
F98
F87
R90
F57
R180
F3
L90
F60
N3
F94
N3
E2
S2
W2
L90
F26
R90
W1
F58
S5
F22
N3
F95
N1
W1
F11
R180
S3
R270
N1
N1
E4
S3
F81
W2
S2
L90
S4
R90
S1
E1
L90
S4
E2
F72
S1
W1
F73
W4
L180
S1
W5
S5
R90
E3
N3
F28
N1
F84
R90
E5
F13
W5
L90
F11
E4
F63
S1
S2
L90
N3
S5
F20
W1
S4
W1
S2
F67
N1
R180
F11
E1
R90
S2
R90
F38
S5
F27
S5
W3
S3
L90
N2
W2
S1
N4
R90
E1
F16
L90
E4
N1
L90
F89
E5
F90
E1
L180
N1
E3
S2
F58
S3
F53
R90
F100
W1
F53
W5
L90
W3
N4
F54
R180
S2
E2
F63
L90
S4
F40
F90
N3
F83
E5
F35
W4
W5
S3
E2
S2
S1
F12
L90
S1
F41
R90
S3
R90
F66
S2
F34
N5
R90
E5
R90
F25
N5
R90
W1
S2
S2
R90
E3
R90
F95
N3
W2
S4
R90
E3
L90
E3
R90
N2
F84
L90
N2
R90
S1
L90
F93
L90
F60
S4
F85
S2
F84
R180
W5
N4
W5
R270
S2
E2
L180
W4
R180
W5
F56
E1
F45
W4
R90
L180
S1
W4
S5
F87
R180
S2
F76
R90
F76
S1
E4
F6
S1
E2
F47
S3
W2
F16
F75
E3
F75
E4
R90
N5
W3
F1
S1
F8
E2
F64
R90
W4
S5
R90
N5
R90
E2
N1
E1
L180
F31
L180
E5
L90
N3
R90
F77
E3
F65
E4
R90
W1
N3
E3
F4
R90
E3
N4
F28
R180
N2
L90
S2
L90
N1
W1
L180
E4
F51
W4
F9
S4
R90
W5
S4
R90
E3
W2
F44
R90
E1
L180
S4
F93
S2
F58
R90
F80
L90
E2
F20
R90
F19
S4
F22
W1
S2
F62
N2
E5
F21
L90
F16
W2
F58
E2
F54
N1
F83
N3
E2
F62
S3
L90
E3
L270
F29
N5
L90
S2
F19
E5
R180
F87
R180
S2
F22
W3
S5
F35
E3
N5
R180
E2
R270
N3
F5
L90
W3
S1
L90
S2
R270
N5
L180
F79
N3
F82
N3
F73
N4
F57
L90
W3
F26
N4
E5
N4
F48
R90
F62
R90
F36
E1
F76
R90
N3
F83
E5
L90
S1
F1
E1
L90
F67
W3
L90
F42
E4
S2
L180
F89
N1
E4
S4
W1
S2
L90
F91
R90
F78
N5
F29
W2
R90
W1
R90
E2
F40
E2
F76
R90
E4
L90
W3
S1
W3
N4
F81
W4
F22
N1
W1
F47
E3
R90
N5
W4
L90
F44
L90
F58
S1
R90
E3
F91
N1
W4
N5
L90
F60
F8
S4
F17
E5
N2
L90
F37
L180
W4
L180
W4
F93
S5
F71
R90
N5
E3
F20
R90
N5
E5
R90
S2
R180
S5
L90
F26
E4
F49
E5
S4
E4
N1
L180
F33
E1
L90
S4
E5
N5
L180
F60
S1
F53
W1
F34
E2
N3
E2
S5
F61
F32
F18
L90
W4
N4
R90
E4
L90
F26
S3
W4
S5
E5
S2
R90
S5
R90
E2
R90
N2
S3
L90
S3
S2
L90
E5
F31
R180
E2
F42
R90
W4
L90
E5
S3
F79
R90
F76
W1
F9
N1
F91
R90
N3
F32
L180
N4
W2
F18
N5
L90
S4
L180
F22
E4
R180
S2
L90
W5
F31
L180
E1
F15
W1
S5
E4
F56
L90
F7
S3
E4
F13
R90
S4
F78
R180
E3
N1
R90
W4
F29
S4
L90
S4
W4
L180
F6
F84
E4
L90
F29
E2
F75
E3
N4
F33
L270
S4
F62
E2
N4
F50
E1
E2
L270
F100
R180
S2
F89
N2
W3
F5
E5
N2
F75
S3
F2
S4
W3
F33
S5
R90
E5
F20
S1
L90
S4
F10
W5
S2
L90
E3
E1
S3
F41
E5
L90
F28
S2
F40
S2
S4
N1
F15
W5
W2
R270
E5
R90
E4
F50
R90
S4
E1
N4
F73
L90
W2
L270
E4
S1
F30
S3
W4
L180
W4
F6
R90
F43
N5
W3
W3
R90
S2
R90
F50
W2
S4
L90
L180
W1
L270
E3
N2
E2
F48
W1
L90
N4
L90
E3
F86
L180
F10
S1
F36
S4
F33
N5
L90
F14
S2
E4
N4
R90
S2
W2
N2
L180
N4
R90
L90
E5
S1
F82
W1
S1
F19
E3
F19
S2
W1
F18
L90
F43
W1
N2
E3
L90
S1
F75
L90
E4
F80
N5
F15
F53
N4
E5
L90
E1
W5
R270
F17
F63
W2
F20
E1
S3
F19
E1
F99
S1
W1
F23
L180
S3
W4
F76`
