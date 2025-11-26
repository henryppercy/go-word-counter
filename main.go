package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	flag.BoolVar(&opts.ShowBytes, "c", false, "Used to toggle whether or not to show the character count")
	flag.BoolVar(&opts.ShowWords, "w", false, "Used to toggle whether or not to show the word count")
	flag.BoolVar(&opts.ShowLines, "l", false, "Used to toggle whether or not to show the line count")

	flag.Parse()

	log.SetFlags(0)

	totals := Counts{}

	fileNames := flag.Args()
	didError := false

	for _, f := range fileNames {
		counts, err := CountFile(f)

		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}

		totals = totals.Add(counts)

		counts.Print(os.Stdout, opts, f)
	}

	if len(fileNames) == 0 {
		counts := GetCount(os.Stdin)
		counts.Print(os.Stdout, opts)
	}

	if len(fileNames) > 1 {
		totals.Print(os.Stdout, opts, "total")
	}

	if didError {
		os.Exit(1)
	}
}
