package cmd

import "strings"

type WordCounter struct {
	Text string
}

func NewWordCounter(text string) *WordCounter {
	return &WordCounter{Text: text}
}

func (wc *WordCounter) CountWords() int {
	words := strings.Fields((wc.Text))

	return len(words)
}
