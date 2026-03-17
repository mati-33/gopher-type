package modes

import (
	"math/rand"
)

type plDiacriticsMode struct {
	letters []rune
	min     int
	max     int
}

func (p plDiacriticsMode) Generate(wordCount int) []rune {
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

func (p plDiacriticsMode) Name() string {
	return "polish diacritics"
}

func (p plDiacriticsMode) genWord() []rune {
	length := rand.Intn(p.max-p.min) + p.min
	word := make([]rune, 0, length)

	for len(word) < length {
		word = append(word, p.letters[rand.Intn(len(p.letters))])
	}

	return word
}

func newPlDiacriticsMode() Mode {
	return plDiacriticsMode{
		letters: []rune("ąćęłśńóżź"),
		min:     3,
		max:     8,
	}
}
