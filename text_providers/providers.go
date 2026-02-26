package textproviders

import (
	_ "embed"
	"math/rand"
	"strings"
	"unicode/utf8"
)

//go:embed eng_top_100_4_letter_words.txt
var EngTop100_4LetterWords string

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

func (p WordArrayProvider) Provide(maxLen int) []rune {
	l := len(p.words)
	b := strings.Builder{}

	for {
		w := p.words[rand.Intn(l)]

		if b.Len()+utf8.RuneCountInString(w) > maxLen {
			break
		}

		b.WriteString(w)
		b.WriteString(" ")
	}

	return []rune(b.String()[:b.Len()-1])
}
