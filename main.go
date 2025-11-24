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
		counts, err := CountFile(f)

		if err != nil {
			fmt.Fprintln(os.Stderr, "counter:", err)
			didError = true
			continue
		}

		total += counts.Words

		fmt.Println(counts.Lines ,counts.Words, counts.Bytes, f)
	}

	if len(fileNames) == 0 {
		counts := GetCount(os.Stdin)
		fmt.Println(counts.Lines ,counts.Words, counts.Bytes)
	}

	if len(fileNames) > 1 {
		fmt.Println("total", total)
	}

	if didError {
		os.Exit(1)
	}
}
