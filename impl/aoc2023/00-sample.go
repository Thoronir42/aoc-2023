package aoc2023

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type SampleTask struct {
}

func (task SampleTask) Pt1(in io.Reader) string {
	var sin = bufio.NewScanner(in)

	var out = ""
	for sin.Scan() {
		out += sin.Text()
	}
	return out
}

func (task SampleTask) Pt2(in io.Reader) string {
	var sum = 0

	var sin = bufio.NewScanner(in)
	for sin.Scan() {
		var line = sin.Text()
		var num, err = strconv.Atoi(line)
		if err != nil {
			continue
		}
		sum += num
	}

	return fmt.Sprintf("%d", sum)
}

func Create00() SampleTask {
	return SampleTask{}
}
