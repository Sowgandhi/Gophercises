package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("C:/Users/Rishi/Desktop/Gophercises/Exercise_1/problems.csv")
	if err != nil {
		fmt.Printf("The file cannot be opened %s\n", err)
		os.Exit(1)
	}
	r := csv.NewReader(file)
  count := 0
	total := 0
	for {
		var ans string
		lines, err := r.Read()
		if err == io.EOF {
			//fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Println("Error")
		}

		fmt.Printf("Question: %s Answer:?\n", lines[0])
		fmt.Scan(&ans)
		if ans == lines[1] {
			count++
		}
		total++
	}
	fmt.Printf("You scored %d out of %d questions", count, total)

}
