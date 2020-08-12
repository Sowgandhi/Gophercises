package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var set int = 0 //used to switch between go routines
var count int = 0
var total int = 0

func main() {

	var path string
	fmt.Printf("Enter the path to the csv file:")
	fmt.Scan(&path)
	pathptr := flag.String("Path", path, "Path to the csv file")
	flag.Parse()
	file, err := os.Open(*pathptr)
	if err != nil {
		fmt.Printf("The file cannot be opened %s\n", err)
		os.Exit(1)
	}
	record := csv.NewReader(file)
	limit := flag.Float64("duration", 30, "Default duration of the quiz")
	flag.Parse()

	now := time.Now()

	for {

		set = 0
		go waiting(now, *limit)
		var ans string
		lines, err := record.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Printf("Question: %s Answer:?\n", lines[0])
		total++
		fmt.Scan(&ans)
		set = 1
		if ans == lines[1] {
			count++
		}

	}
	fmt.Printf("You scored %d out of %d questions", count, total)
}
func waiting(then time.Time, limit float64) {
	for set != 1 {
		now := time.Now()
		if int(now.Sub(then).Seconds()) >= int(limit) {
			fmt.Printf("Time Up!\n")
			fmt.Printf("You scored %d out of %d questions\n", count, total)
			os.Exit(1)

		}
	}
}
