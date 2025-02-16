package main

import (
	"bufio"

	"log"
	"os"
)

func main() {
	// Input file handling
	inputFile, err := os.Open("write")
	if err != nil {
		log.Printf("Error opening input file: %v", err)
		return
	}

	defer inputFile.Close()

	// Output file handling
	outputFile, err := os.Create("display")
	if err != nil {
		log.Printf("Error creating output file: %v", err)
		return
	}

	defer outputFile.Close()

	// Process file line by line
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		modified := line

		// Apply external modifications
		modified = edits.HandleHex(modifed)
		modified = edits.HandleBinary(modified)
		modified = edits.Handlecadifications(modified)
		modified = edits.HandlePunctuation(modified)
		modified = edits.HandleApostrophes(modified)
		modified = edits.HandleArticles(modified)

		// Write modified line
		_, err = outputFile.WriteString(modified + "\n")
		if err != nil {
			log.Printf("Error writing to output file: %v", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning file: %v", err)
		return
	}
}
