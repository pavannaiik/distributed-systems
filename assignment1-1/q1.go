package cos418_hw1_1

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
//
//	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
//
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuations and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"
	// Read the file content
	data, err := os.ReadFile(path)
	//checkError(err)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	text := string(data)
	text = strings.ToLower(text)

	reg, err := regexp.Compile("[^0-9a-zA-Z]+")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Split the cleaned text into words
	words := strings.Fields(text)

	// Map to store word counts
	wordCounts := make(map[string]int)

	// Count the words that meet the charThreshold
	for _, word := range words {
		cleanedWord := reg.ReplaceAllString(word, "")
		if len(cleanedWord) >= charThreshold {
			wordCounts[cleanedWord]++
		}
	}

	// Convert the map to a slice of WordCount
	var wordCountList []WordCount
	for word, count := range wordCounts {
		wordCountList = append(wordCountList, WordCount{Word: word, Count: count})
	}

	// Sort the word counts
	sortWordCounts(wordCountList)

	// Return the top numWords results
	if numWords > len(wordCountList) {
		numWords = len(wordCountList)
	}
	return wordCountList[:numWords]
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}
