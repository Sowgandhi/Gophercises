package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

/*used to switch between goroutines*/
type globalVariable struct {
	set bool
}

var instance globalVariable

func checkPath(path string) string {
	_, file := filepath.Split(path)
	if file != "problems.csv" {
		path = filepath.Join(path, "problems.csv")
	}
	return path
}

func main() {

	instance.set = false
	var path string
	fmt.Printf("Enter the path to the csv file:")
	fmt.Scan(&path)
	path = checkPath(path)
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
	count := 0
	total := 0
	now := time.Now()

	for {
		instance.set = false
		go waiting(now, *limit, count, total)
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
		instance.set = true
		if ans == lines[1] {
			count++
		}

	}
	fmt.Printf("You scored %d out of %d questions", count, total)
}
func waiting(then time.Time, limit float64, count int, total int) {
	for instance.set != true {
		now := time.Now()
		if int(now.Sub(then).Seconds()) >= int(limit) {
			fmt.Printf("Time Up!\n")
			fmt.Printf("You scored %d out of %d questions\n", count, total+1)
			os.Exit(1)
		}
	}
}
