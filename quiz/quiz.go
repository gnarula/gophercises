package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var csvPtr = flag.String("csv", "problems.csv", "a csv file in the format of question,answer'")
var limitPtr = flag.Int("limit", 30, "the time limit for the quiz in seconds")

func parseCsv(filename string) [][]string {
	file, err := os.Open(*csvPtr)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}

func calculateScore(records [][]string) int {
	count := 0
	for i, record := range records {
		var input string
		fmt.Printf("Problem #%v: %s = ", i+1, record[0])
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)
		if input == record[1] {
			count++
		}
	}
	return count
}

func main() {
	flag.Parse()

	records := parseCsv(*csvPtr)
	score := calculateScore(records)
	fmt.Printf("You scored %v out of %v\n", score, len(records))
}
