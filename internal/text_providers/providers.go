package textproviders

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed words/eng_1k.txt
var Eng1k string

type WordArrayProvider struct {
	words []string
}

func NewWordArrayProvider(words []string) WordArrayProvider {
	return WordArrayProvider{words: words}
}

func NewWordArrayProviderFromTxtFile(fileContent string) WordArrayProvider {
	words := strings.Fields(fileContent)
	return WordArrayProvider{words: words}
}

func (p WordArrayProvider) Provide(wordCount int) []rune {
	l := len(p.words)
	b := strings.Builder{}

	for range wordCount {
		w := p.words[rand.Intn(l)]
		b.WriteString(w)
		b.WriteString(" ")
	}

	return []rune(b.String()[:b.Len()-1])
}
