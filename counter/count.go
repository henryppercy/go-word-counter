package counter

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/henryppercy/counter/display"
)

type Counts struct {
	bytes int
	words int
	lines int
}

func (c Counts) Add(other Counts) Counts {
	c.lines += other.lines
	c.words += other.words
	c.bytes += other.bytes
	return c
}

func (c Counts) Print(w io.Writer, opts display.Options, suffixes ...string) {
	stats := []string{}

	if opts.ShouldShowLines() {
		stats = append(stats, strconv.Itoa(c.lines))
	}

	if opts.ShouldShowWords() {
		stats = append(stats, strconv.Itoa(c.words))
	}

	if opts.ShouldShowBytes() {
		stats = append(stats, strconv.Itoa(c.bytes))
	}

	line := strings.Join(stats, "\t") + "\t"

	fmt.Fprint(w, line)

	suffixStr := strings.Join(suffixes, " ")
	if suffixStr != "" {
		fmt.Fprintf(w, " %s", suffixStr)
	}

	fmt.Fprint(w, "\n")
}

func GetCountSinglePass(f io.Reader) Counts {
	res := Counts{}

	isInsideWord := false
	reader := bufio.NewReader(f)

	for {
		r, size, err := reader.ReadRune()
		if err != nil {
			break
		}

		res.bytes += size
		if r == '\n' {
			res.lines++
		}

		isSpace := unicode.IsSpace(r)

		if !isSpace && !isInsideWord {
			res.words++
		}

		isInsideWord = !isSpace
	}

	return res
}

func GetCount(r io.Reader) Counts {
	bytesReader, bytesWriter := io.Pipe()
	wordsReader, wordsWriter := io.Pipe()
	linesReader, linesWriter := io.Pipe()

	w := io.MultiWriter(bytesWriter, wordsWriter, linesWriter)

	chBytes := make(chan int)
	chWords := make(chan int)
	chLines := make(chan int)

	go func() {
		defer close(chBytes)
		chBytes <- CountBytes(bytesReader)
	}()

	go func() {
		defer close(chWords)
		chWords <- CountWords(wordsReader)
	}()

	go func() {
		defer close(chLines)
		chLines <- CountLines(linesReader)
	}()

	io.Copy(w, r)
	bytesWriter.Close()
	wordsWriter.Close()
	linesWriter.Close()

	byteCount := <-chBytes
	wordCount := <-chWords
	lineCount := <-chLines

	return Counts{
		bytes: byteCount,
		words: wordCount,
		lines: lineCount,
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
