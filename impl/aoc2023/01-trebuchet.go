package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strings"
)

type TrebuchetTask struct {
}

func (task TrebuchetTask) Pt1(in io.Reader) string {
	var sin = bufio.NewScanner(in)

	var numbers = loadNumbers(sin, false)

	var sum = 0
	for _, n := range numbers {
		sum += n
	}

	return fmt.Sprintf("%d", sum)
}

func (task TrebuchetTask) Pt2(in io.Reader) string {
	var sin = bufio.NewScanner(in)

	var numbers = loadNumbers(sin, true)

	var sum = 0
	for _, n := range numbers {
		sum += n
	}

	return fmt.Sprintf("%d", sum)
}

func loadNumbers(sin *bufio.Scanner, findWordCiphers bool) []int {
	var numbers = make([]int, 0)

	for sin.Scan() {
		var line = sin.Text()
		if len(line) < 1 {
			continue
		}

		if findWordCiphers {
			line = replaceWordCiphers(line)
		}
		var num = lineToInt(line)
		numbers = append(numbers, num)

		fmt.Println("---")
	}

	return numbers
}

var nums = [...]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func replaceWordCiphers(line string) string {
	var orig = line

	// for i, cipher := range nums {
	// naive approach; transforms eightwothree => eigh23
	// line = strings.ReplaceAll(line, cipher, fmt.Sprint(i))
	// }

	var replaced = "_"
	for replaced != "" {
		line, replaced = tryReplace(line)
		fmt.Println(replaced)
	}

	fmt.Println(orig, line)
	return line
}
func tryReplace(line string) (string, string) {
	var iMin int = math.MaxInt
	var nMin = -1

	for n, cipher := range nums {
		var i = strings.Index(line, cipher)
		if i == -1 {
			continue
		}
		if i < int(iMin) {
			iMin = i
			nMin = n
		}
	}

	if iMin == math.MaxInt {
		return line, ""
	}

	// one => 1e to leave letters for overlapping ciphers
	var replacee = fmt.Sprintf("%d%c", nMin, nums[nMin][len(nums[nMin])-1])
	return strings.Replace(line, nums[nMin], replacee, 1), replacee
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

	fmt.Printf("[%c%c] %s\n", first, last, line)

	return ((int)(first-'0'))*10 + (int)(last-'0')
}

func charIsNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
