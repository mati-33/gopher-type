package modes

import (
	"embed"
	"math/rand"
	"strings"
)

//go:embed words/*.txt
var wordsDir embed.FS

type fileMode struct {
	words []string
}

func (p fileMode) Generate(wordCount int) []rune {
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

type en1kMode struct {
	fileMode
}

func (p en1kMode) Name() string {
	return "english"
}

func newEn1kMode() Mode {
	file, _ := wordsDir.ReadFile("words/en_1k.txt")
	return en1kMode{
		fileMode{words: strings.Fields(string(file))},
	}
}

type pl2kMode struct {
	fileMode
}

func (p pl2kMode) Name() string {
	return "polish"
}

func newPl2kMode() Mode {
	file, _ := wordsDir.ReadFile("words/pl_2k.txt")
	return pl2kMode{
		fileMode{words: strings.Fields(string(file))},
	}
}
