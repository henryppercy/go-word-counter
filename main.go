package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
)

type DisplayOptions struct {
	ShowBytes bool
	ShowWords bool
	ShowLines bool
}

func (d DisplayOptions) ShouldShowBytes() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines {
		return true
	}

	return d.ShowBytes
}

func (d DisplayOptions) ShouldShowWords() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines {
		return true
	}

	return d.ShowWords
}

func (d DisplayOptions) ShouldShowLines() bool {
	if !d.ShowBytes && !d.ShowWords && !d.ShowLines {
		return true
	}

	return d.ShowLines
}

func main() {
	opts := DisplayOptions{}
	header := false

	flag.BoolVar(&header, "header", false, "Show column titles")
	flag.BoolVar(&opts.ShowBytes, "c", false, "Show the character count")
	flag.BoolVar(&opts.ShowWords, "w", false, "Show the word count")
	flag.BoolVar(&opts.ShowLines, "l", false, "Show the line count")

	flag.Parse()

	log.SetFlags(0)

	wr := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := Counts{}

	fileNames := flag.Args()
	didError := false

	if header {
		headers := []string{}

		if opts.ShouldShowLines() {
			headers = append(headers, "lines")
		}

		if opts.ShouldShowWords() {
			headers = append(headers, "words")
		}

		if opts.ShouldShowBytes() {
			headers = append(headers, "bytes")
		}

		line := strings.Join(headers, "\t") + "\t"

		fmt.Fprintln(wr, line)
	}

	for _, f := range fileNames {
		counts, err := CountFile(f)

		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}

		totals = totals.Add(counts)

		counts.Print(wr, opts, f)
	}

	if len(fileNames) == 0 {
		counts := GetCount(os.Stdin)
		counts.Print(wr, opts)
	}

	if len(fileNames) > 1 {
		totals.Print(wr, opts, "total")
	}

	wr.Flush()
	if didError {
		os.Exit(1)
	}
}
