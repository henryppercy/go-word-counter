package main

import (
	"os"
	"fmt"
)

func main () {
	data, _ := os.ReadFile("./words.txt")

	wordCount := countWords(data);
	
	fmt.Println(wordCount)
}

func countWords(data []byte) int {
	wordCount := 0

	for _, x := range data {
		if x == ' ' {
			wordCount++
		}
	}

	wordCount++

	return wordCount
}
