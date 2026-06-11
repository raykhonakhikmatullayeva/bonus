package word

import (
	"cmp"
	"sort"
)

type WordFreq struct {
	Word  string
	Count int
}

func WordCount(words []string) []WordFreq {
	if len(words) == 0 {
		return nil
	}
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	c := make([]WordFreq, 0, len(counts))
	for word, count := range counts {
		c = append(c, WordFreq{word, count})
	}
	sort.Slice(c, func(a, b int) bool {
		return cmp.Compare(c[b].Word, c[a].Word) < 0
	})
	return c
}
