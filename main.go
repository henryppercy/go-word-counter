package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	totals := Counts{}

	fileNames := os.Args[1:]
	didError := false

	for _, f := range fileNames {
		counts, err := CountFile(f)

		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}

		totals = Counts{
			Bytes: totals.Bytes + counts.Bytes,
			Words: totals.Words + counts.Words,
			Lines: totals.Lines + counts.Lines,
		}

		counts.Print(os.Stdout, f)
	}

	if len(fileNames) == 0 {
		counts := GetCount(os.Stdin)
		counts.Print(os.Stdout, "")
	}

	if len(fileNames) > 1 {
		totals.Print(os.Stdout, "total")
	}

	if didError {
		os.Exit(1)
	}
}
