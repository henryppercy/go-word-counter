package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"text/tabwriter"

	"github.com/henryppercy/counter/counter"
	"github.com/henryppercy/counter/display"
)

type FileCountsResults struct {
	counts   counter.Counts
	fileName string
	err      error
	index    int
}

func main() {
	args := display.NewOptionsArgs{}

	header := false

	flag.BoolVar(&header, "header", false, "Show column titles")
	flag.BoolVar(&args.ShowBytes, "c", false, "Show the character count")
	flag.BoolVar(&args.ShowWords, "w", false, "Show the word count")
	flag.BoolVar(&args.ShowLines, "l", false, "Show the line count")

	flag.Parse()

	opts := display.NewOptions(args)

	log.SetFlags(0)

	wr := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)

	totals := counter.Counts{}

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

	ch := CountFiles(fileNames)

	results := make([]FileCountsResults, len(fileNames))

	for res := range ch {
		results[res.index] = res
	}

	for _, res := range results {
		if res.err != nil {
			didError = true
			fmt.Fprintln(os.Stderr, "counter:", res.err)
			continue
		}

		totals = totals.Add(res.counts)
		res.counts.Print(wr, opts, res.fileName)
	}

	if len(fileNames) == 0 {
		counts := counter.GetCount(os.Stdin)
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

func CountFiles(filenames []string) <-chan FileCountsResults {
	ch := make(chan FileCountsResults)

	wg := sync.WaitGroup{}
	wg.Add(len(filenames))

	for i, filename := range filenames {
		go func() {
			defer wg.Done()

			res, err := counter.CountFile(filename)

			ch <- FileCountsResults{
				counts:   res,
				fileName: filename,
				err:      err,
				index:    i,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}
