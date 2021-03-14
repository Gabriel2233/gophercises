package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var points = 0
var problemCount = 1

type problem struct {
	question string
	answer   string
}

func promptQuestionToUser(p problem) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Problem #%d: what is %v? ", problemCount, p.question)

	for scanner.Scan() {
		text := scanner.Text()

		if text == p.answer {
			points++
		}
		problemCount++
		break
	}
}

func NewTimer(seconds int) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func() {
		<-timer.C
		fmt.Printf("\nTime is over! you got %v points\n", points)
		os.Exit(1)
	}()

	return timer
}

func main() {

	var csvFilename string
	var duration int

	flag.StringVar(&csvFilename, "csv", "problems.csv", "a csv file with the 'question,answer' format")
	flag.IntVar(&duration, "d", 30, "the duration in which the quiz will still be running")
	flag.Parse()

	f, err := os.Open(csvFilename)

	if err != nil {
		fmt.Printf("couldn't open the CSV file, %s\n", csvFilename)
		fmt.Printf("error -> %s\n", err)
		os.Exit(1)
	}

	r := csv.NewReader(f)

	timer := NewTimer(duration)
	defer timer.Stop()

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("couldn't parse the CSV file, %s\n", csvFilename)
			fmt.Printf("error -> %s\n", err)
			os.Exit(1)
		}

		problem := problem{
			question: record[0],
			answer:   record[1],
		}

		promptQuestionToUser(problem)
	}

	fmt.Printf("you've got %v out of %v\n", points, problemCount-1)
}
