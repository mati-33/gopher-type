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

type en1kProvider struct {
	fileProvider
}

func (p en1kProvider) Name() string {
	return "english"
}

func (p en1kProvider) Preview() string {
	return "rope by paragraph sound small match country best thought agree chord came famous car describe"
}

func newEn1kProvider() Provider {
	file, _ := wordsDir.ReadFile("words/en_1k.txt")
	return en1kProvider{
		fileProvider{words: strings.Fields(string(file))},
	}
}

type pl2kProvider struct {
	fileProvider
}

func (p pl2kProvider) Name() string {
	return "polish"
}

func (p pl2kProvider) Preview() string {
	return "miejscowość papier znak lęk narzędzie równocześnie dawny cienki czerwony usłyszeć padać przyprawa odkąd spokojnie"
}

func newPl2kProvider() Provider {
	file, _ := wordsDir.ReadFile("words/pl_2k.txt")
	return pl2kProvider{
		fileProvider{words: strings.Fields(string(file))},
	}
}
