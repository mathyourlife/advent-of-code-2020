# advent-of-code-2020
https://adventofcode.com/2020

This is the first year I decided to participate in Advent of Code.  Previous
work to this was playing around with https://projecteuler.net/ which
I started in 2015 in order to get some more exercise with Go which at that
point I had used for a year or two.

Five years later, some colleagues started a few private leaderboards
for Advent of Code, and that seemed like a good excuse to
get the "just build something" itch out.  As for style, so far, I've stuck
with a single threaded approach.  As with Project Euler, puzzel inputs
are often meant to test the scalability of your algorithm, so spreading
the work across a couple of CPU's isn't going to get you much further.

Here is a summary of some of the skills required for each day.  Days omitted
are designated as general purpose puzzles.

Scalability: Day(s) 8, 10, 13
Recursion (yuck): Day(s) 7
Matrix/Grid: Day(s) 3, 11
Bit manipulation: Day(s) 5
Coordinate Translation/Rotation: Day(s) 12

The `main` branch here will be delayed by a few days to help keep the
overall AoC competition honest.

## When a new problem drops

1. Write the "puzzle input" contents to `./data/dayNN`
2. Run the commands
```
DAY=15
cp solver/day_template.go "solver/day${DAY}.go"
cp solver/day_template_test.go "solver/day${DAY}_test.go"
sed -i "s/DayTEMPLATE/Day$DAY/g" "solver/day${DAY}.go"
sed -i "s/DayTEMPLATE/Day$DAY/g" "solver/day${DAY}_test.go"
```
3. Add a case to the switch statement in [solver/solver.go](solver/solver.go)

## Running Tests

Run the full test suite with:

```
go test -v github.com/mathyourlife/advent-of-code-2020/...
```

Run a specific day's tests with:

```
go test -v -run TestDay12 github.com/mathyourlife/advent-of-code-2020/...
```

## Running a day's solution

To run a specific day, supply the input puzzle file.  The application will
select the corresponding solver based on the day number in the filename.

```
go run github.com/mathyourlife/advent-of-code-2020 data/dayNN
```

For example, to run the day 12 solution provide the [data/day12](data/day12)
file as an argument.  The [main.go](main.go) will parse out the `12` and
use the `solver.NewSolver(day int)` function to return a `Solver` for
that day's input.  All days implement the `Solver` interface:

```
type Solver interface {
	SolvePart1(content string)
	SolvePart2(content string)
}
```

Where `content` provided is expected to have all leading and trailing
spaces trimmed, and the results of the solver are printed to the console.

## How to follow along with the solution

Since the problems have tended to provide a few examples inputs and solutions
to go along with the puzzle, a test driven development (TDD) workflow
made sense to follow.  I don't expect many people to be digging through the
code so it's not well documented.  In some spots I left notes for
myself or commented out log lines in case I feel the need to revisit
or improve the logic.  But, the process is generally to start with the
puzzle such as puzzle 2 where you're checking for valid passwords
according to policies with inputs of `1-3 b: cdefg`.

The rough to day 2 is:

1. Split input content into lines
2. Parse each line to extract data
3. Check if the line contains a valid password
4. Tally up how many valid passwords

From a TDD approach, that translated to creating tests in the following order.

1. (If input parsing was more than splitting on `\n`, there'd be a test here)
2. `TestDay02Parse` - ensure line parsing is correct
3. `TestDay02checkPosition` - ensure that policy checks for the part 1 approach work
4. `TestDay02Part1` - ensure that a full tally in part 1 returns correctly
5. `TestDay02checkFrequency` - ensure that policy checks for the part 2 approach work
6. (If tallying was significantly different in part 2, there'd be a test here)
