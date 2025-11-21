package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	total := 0
	fileNames := os.Args[1:]
	didError := false

	for _, f := range fileNames {
		wordCount, err := CountWordsInFile(f)

		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}

		total += wordCount

		fmt.Println(f, wordCount)
	}

	if len(fileNames) == 0 {
		wordCount := CountWords(os.Stdin)
		fmt.Println(wordCount)
	}

	if len(fileNames) > 1 {
		fmt.Println("total", total)
	}

	if didError {
		os.Exit(1)
	}
}

func CountWordsInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	count := CountWords(file)

	return count, nil
}

func CountWords(file io.Reader) int {
	wordCount := 0

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}
