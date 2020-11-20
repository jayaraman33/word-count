package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(s string) Frequency {
	freq := make(Frequency)
	result := regexp.MustCompile(`([a-z]+'?[a-z]+|[0-9])+`)
	words := result.FindAllString(strings.ToLower(s), -1)
	for _, word := range words {
		freq[word]++
	}
	return freq
}
