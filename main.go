package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
	"log"
	"time"
)

func timeLimit(ch chan string, limit int) {
	time.Sleep(time.Duration(limit) * time.Second)
	ch <- "Time is out !"
}

func playGame(ch chan string) {
	file, err := os.Open("./problems.csv")
	if err != nil {
        log.Fatal("Impossible to open file:", err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	var line string = ""
	var answer string = ""
	var question string = ""
	var solution string = ""

	for scanner.Scan() {
		line = scanner.Text()
		splittedQuestion := strings.Split(line, ",")
		question = splittedQuestion[0]
		solution = splittedQuestion[1]
		fmt.Println(question, "?")
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")
		if (answer == solution) {
			fmt.Println("Good answer !")
		} else {
			fmt.Println("Wrong answer !")
		}
	}
	ch <- "You answered all questions !"
}

func main() {
	limit := parseArgs()
	if (limit == -1) {
		return
	}
	ch := make(chan string)
	go timeLimit(ch, limit)
	go playGame(ch)
	
	result := <-ch
	fmt.Println(result)
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