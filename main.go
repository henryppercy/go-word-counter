package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./words.txt")

	wordCount := CountWords(data)

	fmt.Println(wordCount)
}

func CountWords(data []byte) int {
	words := strings.Fields(string(data))
	return len(words)
}
