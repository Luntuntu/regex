package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	// Open the text file with the data
	file, err := os.Open("spelling_corrections.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create the CSV file to store the results
	csvFile, err := os.Create("spelling_corrections.csv")
	if err != nil {
		log.Fatalf("Failed to create CSV file: %s", err)
	}
	defer csvFile.Close()

	// Create a CSV writer
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write the header row to the CSV file
	writer.Write([]string{"Incorrect", "Correct"})

	// Read the text file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the current line
		line := scanner.Text()

		// Tokenize by splitting at the "->"
		tokens := strings.Split(line, "->")
		if len(tokens) == 2 {
			// Trim spaces from both parts and write to the CSV file
			writer.Write([]string{strings.TrimSpace(tokens[0]), strings.TrimSpace(tokens[1])})
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	log.Println("CSV file written successfully")
}
