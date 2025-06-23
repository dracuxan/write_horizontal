package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("w_h <text>")
		return
	}

	text := os.Args[1:]
	if len(text) == 0 {
		fmt.Println("Please provide some text to format.")
		return
	}

	maxLen := 0

	for _, word := range text {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}

	for i := range maxLen {
		for j, word := range text {
			if i < len(word) {
				fmt.Printf("%c", word[i])
			} else {
				fmt.Print(" ")
			}
			if j != len(text)-1 {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}
}
