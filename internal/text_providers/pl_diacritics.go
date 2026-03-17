package textproviders

import (
	"math/rand"
)

type plDiacriticsProvider struct {
	letters []rune
	min     int
	max     int
}

func (p plDiacriticsProvider) Provide(wordCount int) []rune {
	avgWordLen := (p.min + p.max) / 2
	ret := make([]rune, 0, wordCount*avgWordLen+wordCount-1)

	for i := range wordCount {
		ret = append(ret, p.genWord()...)

		if i < wordCount-1 {
			ret = append(ret, ' ')
		}
	}

	return ret
}

func (p plDiacriticsProvider) Name() string {
	return "polish diacritics"
}

func (p plDiacriticsProvider) Preview() string {
	return "ńżłć ąąńść śłćąś źąśź ćółżó żźśó żśęę ółńźń ęąóńź żśłć ąółóźó ółę óńęńć ćóżę ćźśąęłó"
}

func (p plDiacriticsProvider) genWord() []rune {
	length := rand.Intn(p.max-p.min) + p.min
	word := make([]rune, 0, length)

	for len(word) < length {
		word = append(word, p.letters[rand.Intn(len(p.letters))])
	}

	return word
}

func newPlDiacriticsProvider() Provider {
	return plDiacriticsProvider{
		letters: []rune("ąćęłśńóżź"),
		min:     3,
		max:     8,
	}
}
