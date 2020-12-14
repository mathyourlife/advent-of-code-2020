package solver

type Solver interface {
	SolvePart1(content string)
	SolvePart2(content string)
}

func NewSolver(day int) Solver {
	switch day {
	case 1:
		return NewDay01()
	case 2:
		return NewDay02()
	case 3:
		return NewDay03()
	case 4:
		return NewDay04()
	case 5:
		return NewDay05()
	case 6:
		return NewDay06()
	case 7:
		return NewDay07()
	case 8:
		return NewDay08()
	case 9:
		return NewDay09()
	case 10:
		return NewDay10()
	case 11:
		return NewDay11()
	case 12:
		return NewDay12()
	case 13:
		return NewDay13()
	case 14:
		return NewDay14()
	}
	return nil
}
