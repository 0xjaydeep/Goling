//go:build ignore

// Solution for Exercise 03 — Word Frequency Counter
// Run with: go run solution.go
package main

import (
	"fmt"
	"sort"
	"strings"
)

func countWords(text string) map[string]int {
	freq := make(map[string]int)
	for _, word := range strings.Fields(text) {
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,!?;:\"'")
		if word == "" {
			continue
		}
		freq[word]++
	}
	return freq
}

func topN(freq map[string]int, n int) []string {
	words := make([]string, 0, len(freq))
	for w := range freq {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if freq[words[i]] != freq[words[j]] {
			return freq[words[i]] > freq[words[j]] // higher count first
		}
		return words[i] < words[j] // tie-break: alphabetical
	})
	if n > len(words) {
		n = len(words)
	}
	return words[:n]
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

	check := []string{"go", "easy", "language", "type"}
	fmt.Println("\nSpot-check counts:")
	for _, w := range check {
		fmt.Printf("  %-10q → %d\n", w, freq[w])
	}
}
