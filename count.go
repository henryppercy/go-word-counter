package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Counts struct {
	Bytes int
	Words int
	Lines int
}

func (c Counts) Add(other Counts) Counts {
	c.Lines += other.Lines
	c.Words += other.Words
	c.Bytes += other.Bytes
	return c
}

func (c Counts) Print(w io.Writer, filenames ...string) {
	fmt.Fprintf(w, "%d %d %d", c.Lines, c.Words, c.Bytes)

	for _, f := range filenames {
		fmt.Fprintf(w, " %s", f)	
	}

	fmt.Fprintf(w, "\n")
}

func GetCount(f io.ReadSeeker) Counts {
	const OffsetStart = 0

	bytes := CountBytes(f)
	f.Seek(OffsetStart, io.SeekStart)

	words := CountWords(f)
	f.Seek(OffsetStart, io.SeekStart)

	lines := CountLines(f)

	return Counts{
		Bytes: bytes,
		Words: words,
		Lines: lines,
	}
}

func CountFile(filename string) (Counts, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Counts{}, err
	}
	defer file.Close()

	counts := GetCount(file)

	return counts, nil
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

func CountLines(r io.Reader) int {
	lineCount := 0

	reader := bufio.NewReader(r)

	for {
		r, _, err := reader.ReadRune()
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
