package aoc2023

import (
	"bufio"
	"fmt"
	"io"
)

type TrebuchetTask struct {
}

func (task TrebuchetTask) Pt1(in io.Reader) string {
	var sin = bufio.NewScanner(in)

	var numbers = make([]int, 0)

	for sin.Scan() {
		var line = sin.Text()
		var num = lineToInt(line)
		numbers = append(numbers, num)
	}

	var sum = 0
	for _, n := range numbers {
		sum += n
	}

	return fmt.Sprintf("%d", sum)
}

func (task TrebuchetTask) Pt2(in io.Reader) string {
	panic("not implemented")
}

// The desired number is concat of first and last digit in the string
func lineToInt(line string) int {
	var first, last byte

	for i := 0; i < len(line); i++ {
		var char = line[i]
		if !charIsNumber(char) {
			continue
		}

		last = char
		if first == 0 {
			first = char
		}
	}

	return ((int)(first-'0'))*10 + (int)(last-'0')
}

func charIsNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
