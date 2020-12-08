package main

func main() {
	days := []DayMaker{
		NewDay01,
	}
	d := days[0]()
	d.SolvePart1()
	d.SolvePart2()
}

type DayMaker func() Day

type Day interface {
	SolvePart1()
	SolvePart2()
}
