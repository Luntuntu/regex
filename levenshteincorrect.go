package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/agnivade/levenshtein"
)

// Create a global corrections map
var corrections = map[string]string{}

// Max Levenshtein distance for allowed corrections
const maxLevenshteinDistance = 3

func main() {
	// Load the corrections from the CSV file
	err := loadCorrectionsFromCSV("spelling_corrections.csv")
	if err != nil {
		fmt.Println("Error loading corrections:", err)
		return
	}

	// Input sentence
	fmt.Println("Enter a sentence:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Before autocorrect
	fmt.Println("Before autocorrect:", input)

	// After autocorrect
	corrected := autocorrect(input)
	fmt.Println("After autocorrect:", corrected)
}

// Load corrections from a CSV file
func loadCorrectionsFromCSV(filePath string) error {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all lines
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Parse the CSV content
	for _, record := range records {
		if len(record) < 2 {
			continue // Skip rows that don't have enough data
		}

		// First column is the misspelled word
		misspelledWord := strings.TrimSpace(record[0])

		// Second column is the correct word
		correctWord := strings.TrimSpace(record[1])

		// Add the misspelled word and correct word to the corrections map
		corrections[misspelledWord] = correctWord
	}

	return nil
}

// Autocorrect function using Levenshtein distance
func autocorrect(input string) string {
	// Split input into words
	words := strings.Fields(input)

	for i, word := range words {
		// Trim punctuation from the word
		cleanWord := strings.Trim(word, ".,!?")

		// Find the closest match based on Levenshtein distance
		closestWord := findClosestWord(cleanWord)

		// If a close match is found, replace the word in the sentence
		if closestWord != cleanWord {
			// Replace the word in the input sentence, preserving punctuation
			words[i] = strings.Replace(word, cleanWord, closestWord, 1)
		}
	}

	// Join the words back into a sentence
	return strings.Join(words, " ")
}

// Find the closest correct word from the corrections map using Levenshtein distance
func findClosestWord(word string) string {
	closestWord := word
	minDistance := len(word) + 1 // Initialize with a high number

	for misspelledWord, correctWord := range corrections {
		// Check the misspelled word first
		if dist := levenshtein.ComputeDistance(word, misspelledWord); dist < minDistance && dist <= maxLevenshteinDistance {
			minDistance = dist
			closestWord = correctWord
		}
	}

	return closestWord
}
