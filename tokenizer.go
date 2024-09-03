package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Specify the file path
	filePath := "input.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define the regex pattern for tokenization
	regexPattern := `\w+|[^\s\w]`

	// Compile the regex
	re := regexp.MustCompile(regexPattern)

	fmt.Println("Tokens:")

	// Process the file line by line
	lineNumber := 1
	for scanner.Scan() {
		line := scanner.Text()
		tokens := re.FindAllString(line, -1)

		// Print the tokens for the current line
		fmt.Printf("Line %d:\n", lineNumber)
		for i, token := range tokens {
			fmt.Printf("  Token %d: %s\n", i+1, token)
		}
		lineNumber++
	}

	// Check for errors in scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
