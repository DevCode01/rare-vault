package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	freq := map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w := strings.ToLower(strings.Trim(scanner.Text(), ".,;:!?\"'()[]{}"))
		if w != "" {
			freq[w]++
		}
	}

	type pair struct {
		word  string
		count int
	}
	pairs := make([]pair, 0, len(freq))
	for w, c := range freq {
		pairs = append(pairs, pair{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].word < pairs[j].word
		}
		return pairs[i].count > pairs[j].count
	})

	total := 0
	for _, p := range pairs {
		total += p.count
	}
	fmt.Printf("Total words: %d\nUnique words: %d\n\n", total, len(pairs))
	fmt.Println("Top 15 words:")
	for i, p := range pairs {
		if i >= 15 {
			break
		}
		fmt.Printf("  %-15s %d\n", p.word, p.count)
	}
}
