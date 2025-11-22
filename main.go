package main

import (
	"fmt"
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
