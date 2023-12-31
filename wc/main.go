package main

import (
	"fmt"
	wc "wc/cmd"
)

func main() {
	text := "Hello, this is a sample text for word counting"
	counter := wc.NewWordCounter(text)
	wordCount := counter.CountWords()

	fmt.Printf("Number of words: %d\n", wordCount)
}
