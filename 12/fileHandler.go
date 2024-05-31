package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Print("Error opening file")
		return
	}

	defer inputFile.Close()

	wordCounts := make(map[string]int)
	lineCount := 0
	specificWord := "file"
	specificWordCount := 0

	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		words := strings.Fields(line)

		for _, word := range words {
			word = strings.ToLower(word)
			wordCounts[word]++
			if word == specificWord {
				specificWordCount++
			}
		}
	}

	outputFile, err := os.Create("output.txt")
	writer := bufio.NewWriter(outputFile)

	writer.WriteString(fmt.Sprintf("Total number of lines: %d\n", lineCount))
	writer.WriteString(fmt.Sprintf("Total occurrences of the word '%s': %d\n", specificWord, specificWordCount))

	writer.Flush()
	fmt.Println("Successfully written to output.txt.")
}
