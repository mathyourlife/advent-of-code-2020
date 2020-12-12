package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	solveDay := os.Args[1]
	solveDayNum, _ := strconv.Atoi(solveDay)
	fmt.Printf("solving Day %s\n", solveDay)
	days := []DayMaker{
		NewDay01,
		NewDay02,
		NewDay03,
		NewDay04,
		NewDay05,
		NewDay06,
		NewDay07,
		NewDay08,
		NewDay09,
		NewDay10,
		NewDay11,
	}
	d := days[solveDayNum-1]()
	d.SolvePart1()
	d.SolvePart2()
}

type DayMaker func() Day

type Day interface {
	SolvePart1()
	SolvePart2()
}
