package solver

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Day14 struct {
	memory map[int]int
}

func NewDay14() Solver {
	return &Day14{
		memory: map[int]int{},
	}
}

func (d *Day14) Reset() {
	d.memory = map[int]int{}
}

func (d *Day14) SolvePart1(content string) {
	d.execPrograms(content)
	sum := 0
	for _, v := range d.memory {
		sum += v
	}
	fmt.Printf("sum of values in memory: %d\n", sum)
}

func (d *Day14) SolvePart2(content string) {
	d.execV2(content)
	sum := 0
	for _, v := range d.memory {
		sum += v
	}
	// 9793721823875 too high
	fmt.Println("sum", sum)
}

func (d *Day14) execV2(content string) {
	programs := strings.Split(content, "mask = ")
	for _, program := range programs {
		if !strings.Contains(program, "mem") {
			continue
		}
		mask, instructions := d.parseProgram(program)
		set, _, floating := d.parseMask(mask)

		for _, instruction := range instructions {
			d.writeV2(instruction[1], instruction[0], set, floating)
		}
	}
}

func (d *Day14) writeV2(val, addr, set int, floating []int) {

	// log.Printf("addr   %4d %10s\n", addr, strconv.FormatInt(int64(addr), 2))
	// log.Printf("set    %4d %10s\n", set, strconv.FormatInt(int64(set), 2))
	maskedAddr := addr | set
	// log.Printf("masked %4d %10s\n", maskedAddr, strconv.FormatInt(int64(maskedAddr), 2))
	index := int64(0)
	// log.Println(floating)

	for {
		binaryStr := strconv.FormatInt(index, 2)
		if len(binaryStr) > len(floating) {
			break
		}

		on := 0
		off := 0
		for pos := 0; pos < len(floating); pos++ {
			if 1<<pos&index > 0 {
				on |= 1 << floating[pos]
			} else {
				off |= 1 << floating[pos]
			}
		}
		// log.Println("ON ", strings.ReplaceAll(fmt.Sprintf("%9s", strconv.FormatInt(int64(on), 2)), " ", "0"),
		// 	"OFF", strings.ReplaceAll(fmt.Sprintf("%9s", strconv.FormatInt(int64(off), 2)), " ", "0"))
		// log.Println("ADDR", strings.ReplaceAll(fmt.Sprintf("%9s", strconv.FormatInt(int64((maskedAddr|on)&(int(math.Pow(2, 36)-1)-off)), 2)), " ", "0"))
		d.memory[(maskedAddr|on)&(int(math.Pow(2, 36)-1)-off)] = val
		index++
	}
}

func (d *Day14) execPrograms(content string) {
	programs := strings.Split(content, "mask = ")
	for _, program := range programs {
		if !strings.Contains(program, "mem") {
			continue
		}
		mask, instructions := d.parseProgram(program)
		set, unset, _ := d.parseMask(mask)
		d.execProgram(instructions, set, unset)
	}
}

func (d *Day14) execProgram(instructions [][]int, set, unset int) {
	for _, instruction := range instructions {
		d.memory[instruction[0]] = d.maskValue(instruction[1], set, unset)
	}
}

func (d *Day14) parseProgram(program string) (string, [][]int) {
	lines := strings.Split(strings.TrimSpace(program), "\n")
	mask := lines[0][len(lines[0])-36 : len(lines[0])]
	re := regexp.MustCompile(`mem\[(\d+)] = (\d+)`)
	instructions := [][]int{}
	for i := 1; i < len(lines); i++ {
		match := re.FindStringSubmatch(lines[i])
		addr, _ := strconv.Atoi(match[1])
		val, _ := strconv.Atoi(match[2])
		instructions = append(instructions, []int{addr, val})
	}
	return mask, instructions
}

func (d *Day14) maskValue(val, set, unset int) int {
	return (val | set) & (int(math.Pow(2, 36)-1) - unset)
}

// return 2 int bit masks set and unset where bits "on" in the set
// int represent bits that should be turned onn, and bits "on" in unset
// represent bits that should be turned off.
func (d *Day14) parseMask(mask string) (set int, unset int, floating []int) {
	floating = []int{}
	for i := 0; i < len(mask); i++ {
		bit := mask[i : i+1]
		set = set << 1
		unset = unset << 1
		switch bit {
		case "X":
			// Do nothing
			floating = append(floating, 35-i)
		case "1":
			set++
		case "0":
			unset++
		}
	}
	return set, unset, floating
}
