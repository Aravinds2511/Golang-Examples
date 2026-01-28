// Package challenge6 contains the solution for Challenge 6.
package main

import (
	// Add any necessary imports here
	"fmt";
    "strings";
	"regexp"
)

// CountWordFrequency takes a string containing multiple words and returns
// a map where each key is a word and the value is the number of times that
// word appears in the string. The comparison is case-insensitive.
//
// Words are defined as sequences of letters and digits.
// All words are converted to lowercase before counting.
// All punctuation, spaces, and other non-alphanumeric characters are ignored.
//
// For example:
// Input: "The quick brown fox jumps over the lazy dog."
// Output: map[string]int{"the": 2, "quick": 1, "brown": 1, "fox": 1, "jumps": 1, "over": 1, "lazy": 1, "dog": 1}

func main() {
	words := "The quick brown fox jumps over the lazy dog."
    fmt.Println(CountWordFrequency(words))
}

func CountWordFrequency(text string) map[string]int {
	// Your implementation here
	wordmap := make(map[string]int)
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	text = re.ReplaceAllString(text, " ")
	words := strings.Fields(text)
	for _,word := range words {
	    wordmap[word]++
	}
	return wordmap
}