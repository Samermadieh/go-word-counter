package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if the filename is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: word_counter <filename>")
		return
	}

	// Get the filename
	filename := os.Args[1]

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a map to store the word count
	wordCount := make(map[string]int)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)

		// Increment the count of each word
		for _, word := range words {
			wordCount[word]++
		}
	}

	// Print the word count
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
