package common

import (
	"errors"
	"fmt"
	"io"
	"os"

	"t42/aoc/impl/aoc2023"
)

type AOCTaskRunner struct {
	DataDir         string
	Year            int64
	UserName        string
	AutoFetchCookie string
}

func (runner AOCTaskRunner) RunTask(task string, part string) (string, error) {
	var t, err = GetTask(task)
	if err != nil {
		return "", err
	}

	var reader, readerErr = runner.createReader(task)
	if readerErr != nil {
		return "", readerErr
	}

	if part == "1" {
		return t.Pt1(reader), nil
	}
	if part == "2" {
		return t.Pt2(reader), nil
	}

	return "", errors.New("Unknown part")
}

func GetTask(task string) (AocTask, error) {
	if task == "00" {
		var t = aoc2023.SampleTask{}

		return t, nil
	}

	return nil, errors.New("Unknown task: " + task)
}

func (runner AOCTaskRunner) createReader(task string) (io.Reader, error) {
	var path = fmt.Sprintf("%s/%d/%s/%s.txt", runner.DataDir, runner.Year, runner.UserName, task)

	// TODO: if file does not exist, try to download it
	//  https://stackoverflow.com/questions/11692860/how-can-i-efficiently-download-a-large-file-using-go

	return os.Open(path)
}
