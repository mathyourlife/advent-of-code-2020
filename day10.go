package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Day10 struct {
}

func NewDay10() Day {
	return &Day10{}
}

func (d *Day10) SolvePart1() {
	fmt.Printf("just used excel\n")
}

func (d *Day10) SolvePart2() {

	lines := strings.Split(Day10Content, "\n")
	vals := make([]int, 0, len(lines)+1)
	vals = append(vals, 0)
	paths := make([]int, len(lines)+1)
	for _, line := range lines {
		i, _ := strconv.Atoi(line)
		vals = append(vals, i)
	}
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

const Day10Content = `152
18
146
22
28
133
114
67
19
37
66
14
90
163
26
149
71
106
46
143
145
12
151
105
58
130
93
49
74
83
129
122
63
134
86
136
166
169
159
3
178
88
103
97
110
53
125
128
9
15
78
1
50
87
56
89
60
139
113
43
36
118
170
96
135
23
144
153
150
142
95
180
35
179
80
13
115
2
171
32
70
6
72
119
29
79
27
47
107
73
162
172
57
40
48
100
64
59
175
104
156
94
77
65`
