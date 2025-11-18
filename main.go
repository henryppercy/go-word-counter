package main

import (
	"os"
	"fmt"
)

func main () {
	data, _ := os.ReadFile("./words.txt")

	wordCount := 0
	// const spaceChar = 32 // decimal value for space 

	for _, x := range data {
		if x == ' ' {           // can use rune instead, single quotes to represent a rune
			wordCount++
		}
	}

	wordCount++

	fmt.Println(wordCount)
}
