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

	wg := sync.WaitGroup{}
	wg.Add(len(fileNames))

	ch := make(chan FileCountsResults)

	for _, f := range fileNames {
		go func() {
			defer wg.Done()

			counts, err := counter.CountFile(f)

			if err != nil {
				fmt.Fprintln(os.Stderr, "counter:", err)
				didError = true
				return
			}

			ch <- FileCountsResults{
				counts: counts,
				fileName: f,
			}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
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
