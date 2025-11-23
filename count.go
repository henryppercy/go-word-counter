package main

import (
	"bufio"
	"io"
	"os"
)

type Counts struct {
	Bytes int
	Words int
	Lines int
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	return Counts{
		Bytes: CountBytes(file),
		Words: CountWords(file),
		Lines: CountLines(file),
	}, nil
}

func CountWords(r io.Reader) int {
	wordCount := 0

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordCount++
	}

	return wordCount
}

func CountLines (r io.Reader) int {
	lineCount := 0

	reader := bufio.NewReader(r)

	for {
		r, _ , err := reader.ReadRune()
		if err != nil {
			break
		}

		if r == '\n' {
			lineCount++
		}
	}

	return lineCount
}

func CountBytes(r io.Reader) int {
	bytesCount, _ := io.Copy(io.Discard, r)

	return int(bytesCount)
}
