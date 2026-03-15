package textproviders

import (
	"embed"
	"math/rand"
	"strings"
)

//go:embed words/*.txt
var wordsDir embed.FS

type fileProvider struct {
	words []string
}

func (p fileProvider) Provide(wordCount int) []rune {
	l := len(p.words)
	b := strings.Builder{}

	for i := range wordCount {
		w := p.words[rand.Intn(l)]
		b.WriteString(w)

		if i < wordCount-1 {
			b.WriteString(" ")
		}
	}

	return []rune(b.String())
}

func newEn1kProvider() Provider {
	file, _ := wordsDir.ReadFile("words/en_1k.txt")
	return fileProvider{words: strings.Fields(string(file))}
}

func newPl2kProvider() Provider {
	file, _ := wordsDir.ReadFile("words/pl_2k.txt")
	return fileProvider{words: strings.Fields(string(file))}
}
