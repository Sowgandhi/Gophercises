package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	filenameptr := flag.String("FileName", "problems.csv", "the path to the csv file")
	limit := flag.Int("duration", 30, "Default duration of the quiz")
	flag.Parse()
	file, err := os.Open(*filenameptr)
	if err != nil {
		fmt.Printf("The file cannot be opened %s\n", err)
		os.Exit(1)
	}
	record := csv.NewReader(file)

	count := 0
	total := 0
	complete := make(chan int)
	go waiting(complete, *limit)

	for {
		lines, err := record.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Printf("Question: %s Answer:?\n", lines[0])
		total++
		answerchan := make(chan string)
		go ans(answerchan)

		select {
		case <-complete:
			fmt.Printf("You scored %d out of %d questions\n", count, total)
			os.Exit(1)
		case answer := <-answerchan:
			if answer == lines[1] {
				count++
			}

		}

	}
}
func ans(answerchan chan string) {
	var answer string
	fmt.Scan(&answer)
	answerchan <- answer
}
func waiting(complete chan int, limit int) {
	time.Sleep(time.Duration(limit) * time.Second)
	complete <- 1
}
