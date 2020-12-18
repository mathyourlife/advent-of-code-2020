package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

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
	contentStr := strings.TrimSpace(string(content))

	solver := solver.NewSolver(day)
	fmt.Println("-------------------- PART 1 --------------------")
	solver.SolvePart1(contentStr)
	fmt.Println("-------------------- PART 2 --------------------")
	solver.Reset()
	solver.SolvePart2(contentStr)
}
