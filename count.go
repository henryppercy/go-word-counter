package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
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

func (c Counts) Print(w io.Writer, opts DisplayOptions, suffixes ...string) {
	xs := []string{}

	if opts.ShouldShowLines() {
		xs = append(xs, strconv.Itoa(c.Lines))
	}

	if opts.ShouldShowWords() {
		xs = append(xs, strconv.Itoa(c.Words))
	}

	if opts.ShouldShowBytes() {
		xs = append(xs, strconv.Itoa(c.Bytes))
	}

	xs = append(xs, suffixes...)

	line := strings.Join(xs, " ")

	fmt.Fprintln(w, line)
}

func GetCount(f io.Reader) Counts {
	res := Counts{}

	isInsideWord := false
	reader := bufio.NewReader(f)

	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}
		
		res.Bytes += size
		if r == '\n' {
			res.Lines++
		}

		isSpace := unicode.IsSpace(r)

		if !isSpace && !isInsideWord {
			res.Words++
		}

		isInsideWord = !isSpace
	}

	return res
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
