package main

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/mathyourlife/advent-of-code-2020/solver"
)

func main() {

	re := regexp.MustCompile(`.*/day(\d+)`)
	inputFile := os.Args[1]
	match := re.FindStringSubmatch(inputFile)
	day, err := strconv.Atoi(match[1])
	if err != nil {
		log.Fatalf("unable to determine what day to solve: %s", inputFile)
	}
	log.Printf("solving day %d", day)

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	solver := solver.NewSolver(day)
	solver.SolvePart1(string(content))
	solver.SolvePart2(string(content))
}
