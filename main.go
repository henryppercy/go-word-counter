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

	if len(os.Args) < 2 {
		log.Fatalln("error: must provide at least one file name")
	}

	total := 0

	fileNames := os.Args[1:]
	for _, f := range fileNames {
		wordCount := CountWordsInFile(f)

		total += wordCount
		
		fmt.Println(f, wordCount)
	}

	if len(fileNames) > 1 {
		fmt.Println("total", total)
	}
}

func CountWordsInFile(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln("failed to read file:", err)
	}
	return CountWords(file)
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
