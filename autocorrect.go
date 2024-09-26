package main

import (
	"fmt"
	"regexp"
)

// Function to correct common misspellings
func autocorrect(input string) string {
	// Define a map of common misspellings to correct words
	corrections := map[*regexp.Regexp]string{
		regexp.MustCompile(`\b(h[aeiou]llo|hrllo|helo)\b`):       "hello",
		regexp.MustCompile(`\b(t[eia]h|teh|t[eoa]e)\b`):          "the",
		regexp.MustCompile(`\b(bey|biye|bai|byee)\b`):            "bye",
		regexp.MustCompile(`\b(tomor+ow|tom+ow|tomrw|tommrw)\b`): "tomorrow",
	}

	for pattern, correctWord := range corrections {
		input = pattern.ReplaceAllString(input, correctWord)
	}

	return input
}

func main() {
	// Sample inputs to test autocorrection
	inputs := []string{
		"hrllo there!",
		"I will see you tomorow.",
		"Can you correct teh text?",
		"Say bey to everyone.",
	}

	for _, input := range inputs {
		corrected := autocorrect(input)
		fmt.Printf("Original: %s\nCorrected: %s\n\n", input, corrected)
	}
}
