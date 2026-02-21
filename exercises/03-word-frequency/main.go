// Exercise 03 — Word Frequency Counter
//
// Problem:
//   Count how often each word appears in a string, then print the top N words
//   sorted by frequency (descending). Break ties alphabetically.
//
// Functions to implement:
//   countWords(text string) map[string]int
//     — split by whitespace, lowercase, strip punctuation, count
//   topN(freq map[string]int, n int) []string
//     — return top n words sorted by count desc, ties broken alphabetically
//
// Hints:
//   strings.Fields(text)         — split by whitespace
//   strings.ToLower(word)        — lowercase
//   strings.Trim(word, ".,!?;:\"'") — strip punctuation
//   sort.Slice(words, fn)        — custom sort
//
// Run your solution:  go run main.go
// See the answer:     go run solution.go
package main

import (
	"fmt"
	"strings"
)

// countWords splits text into words (whitespace-separated), lowercases them,
// strips leading/trailing punctuation, and counts each word's occurrences.
func countWords(text string) map[string]int {
	// TODO: create freq := make(map[string]int)
	// TODO: for _, word := range strings.Fields(text) {
	//         word = strings.ToLower(word)
	//         word = strings.Trim(word, ".,!?;:\"'")
	//         if word == "" { continue }
	//         freq[word]++
	//       }
	// TODO: return freq
	return nil
}

// topN returns the n most frequent words.
// Primary sort: frequency descending. Tie-break: alphabetical ascending.
func topN(freq map[string]int, n int) []string {
	// TODO: build a slice of all keys from freq
	// TODO: sort.Slice with comparator:
	//         if freq[a] != freq[b]: higher freq first
	//         else: alphabetical order
	// TODO: clamp n to len(words) if needed
	// TODO: return words[:n]
	return nil
}

func main() {
	text := `Go is an open source programming language that makes it easy to build
simple reliable and efficient software Go is expressive concise clean and efficient
Its concurrency mechanisms make it easy to write programs that get the most out of
multicore and networked machines while its novel type system enables flexible and
modular program construction Go compiles quickly to machine code yet it still has the
convenience of garbage collection and the power of run time reflection it is a fast
statically typed compiled language that feels like a dynamically typed interpreted language`

	freq := countWords(text)
	fmt.Printf("Unique words: %d\n\n", len(freq))

	fmt.Println("Top 5 words:")
	for i, word := range topN(freq, 5) {
		fmt.Printf("  %d. %-12s %d times\n", i+1, word, freq[word])
	}

	// Spot-check specific words
	check := []string{"go", "easy", "language", "type"}
	fmt.Println("\nSpot-check counts:")
	for _, w := range check {
		fmt.Printf("  %-10q → %d\n", w, freq[w])
	}

	_ = strings.Fields // hint: import already here for you
}
