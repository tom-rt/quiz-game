package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	limit := parseArgs()
	fmt.Println(limit)
}

func parseArgs() int {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a limit.")
		return -1
	}
	arg := os.Args[1]
	split := strings.Split(arg, "=")
	limit, err := strconv.Atoi(split[1])

	if err != nil || len(split) != 2 || split[0] != "-limit"  {
		fmt.Println("Bad argument. Example: \"-limit=2\".")
		return -1
	}
	return limit
}