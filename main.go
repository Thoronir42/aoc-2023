package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"t42/aoc/common"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	var argc = len(os.Args)
	if argc < 2 {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Not enough args"))
		return
	}

	var task = os.Args[1]
	var part = "1"
	if argc > 2 {
		part = os.Args[2]
	}

	var runner, errRunner = createRunnerFromEnv()
	if errRunner != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Failed to create runner: %s", &errRunner))
		return
	}

	var found bool
	part, found = strings.CutPrefix(part, "_")
	if found {
		runner.UserName = "_sample"
	}

	var result, err = runner.RunTask(task, part)

	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Task %s:%s failed", task, part))
		fmt.Fprintln(os.Stderr, fmt.Sprintf("%s", err))
		return
	}

	fmt.Fprintln(os.Stderr, fmt.Sprintf("Task %s:%s completed", task, part))
	fmt.Println(result)
}

func createRunnerFromEnv() (common.AOCTaskRunner, error) {
	var i, err = strconv.ParseInt(os.Getenv("AOC_YEAR"), 10, 64)

	if err != nil {
		return common.AOCTaskRunner{}, err
	}

	return common.AOCTaskRunner{
		DataDir:         os.Getenv("AOC_DATA_DIR"),
		AutoFetchCookie: os.Getenv("AOC_AUTO_FETCH_COOKIE"),
		Year:            i,
		UserName:        os.Getenv("AOC_USER_NAME"),
	}, nil
}
