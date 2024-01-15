package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	//declare word vector out of sentence, split on space
	wordMapped := strings.Fields(s)

	//declare map for outputting words and their counts
	counts := make(map[string]int)

	//looping over each word in the sentence
	//counting the amount of times a word occurs
	for word := range wordMapped {
		counts[wordMapped[word]] += 1
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
