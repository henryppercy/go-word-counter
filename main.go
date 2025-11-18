package main

import (
	"os"
	"fmt"
)

func main () {
	data, _ := os.ReadFile("./words.txt")

	fmt.Println("data:", string(data)) // is a slight performance impact when converting between string and bytes (vica verca), don't go switching back and forth, just as needed
}